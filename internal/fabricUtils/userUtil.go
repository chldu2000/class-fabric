package fabricUtils

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

// RegisterAndEnroll 先采用执行脚本的方式生成用户凭据
func RegisterAndEnroll(userid, orgName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(3 * time.Second)
		cancelFunc()
	}(cancel)
	password := userid + "pwd"
	cmd := "./scripts/createUser.sh " + orgName + " " + userid + " " + password
	err := Command(ctx, cmd)
	if err != nil {
		return err
	}
	return nil
}

func PopulateWallet(wallet *gateway.Wallet, userId string, org string) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network-new",
		"organizations",
		"peerOrganizations",
		"org"+org+".medicine.com",
		"users",
		userId+"@org"+org+".medicine.com",
		"msp",
	)

	log.Println("读取cert……")
	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	log.Println("读取key……")
	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}

	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	log.Println("key: " + string(key))
	log.Println("生成identity……")
	identity := gateway.NewX509Identity("Org"+org+"MSP", string(cert), string(key))
	label := userId + "@" + org

	return wallet.Put(label, identity)
}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}
}

func Command(ctx context.Context, cmd string) error {
	c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	var waitGroup sync.WaitGroup
	// 读取 stderr 和 stdout
	waitGroup.Add(2)
	go read(ctx, &waitGroup, stderr)
	go read(ctx, &waitGroup, stdout)

	err = c.Start()
	// 等待任务结束
	waitGroup.Wait()
	return err
}
