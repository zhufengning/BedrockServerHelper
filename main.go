package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"net/http"
	"encoding/json"
)
type Hitokoto struct {
	Hitokoto string
	From string
	From_who string
}
func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
func printLalala() {
	rand0 := rand.New(rand.NewSource(time.Now().UnixNano() + 954677559))
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	sentences := []string{"袋鼠跳、地震拳、电眼逼人、饿狼前进、二龙戏珠","飞鹤捕虾、飞龙在天、飞天陲、飞天猴巧夺宝盒、飞象踩老鼠、飞鹰展翅","石剑是检验真理的唯一标准!", "你，就是我的Operator吗？", "呐呐呐呐呐呐", "服务器即将爆炸！", "这里是Server酱，老二刺螈了", "林地府邸一根棒，预备，起！", "不定期崩溃不是bug，是本服务器的特性！"}
	for true {
		say := ""
		switch rand0.Int() % 3 {
		case 0: 
			say = sentences[rand1.Int()%len(sentences)]
		case 1:
			resp, err := http.Get("https://yijuzhan.com/api/word.php")
			if err == nil {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					say = string(body)
				} else {
					fmt.Println("me 酱坏掉了：", err)
				}
			} else {
				fmt.Println("me 酱坏掉了:", err)
			}
		case 2:
			resp, err := http.Get("https://v1.hitokoto.cn/?c=a")
			if err == nil {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					hit := &Hitokoto{}
					err := json.Unmarshal(body, &hit)
					if (err == nil ) {
						say = string(hit.Hitokoto)
						say += "——"
						if hit.From_who != "" {
							say += hit.From_who
						} else {
							say += hit.From
						}
					} else {
						fmt.Println("me 酱坏掉了：", err)
					}					
				} else {
					fmt.Println("me 酱坏掉了：", err)
				}
			} else {
				fmt.Println("me 酱坏掉了:", err)
			}
		}
			
		if say != "" {
			fmt.Println("me 酱:", say)
		}
		time.Sleep(1 * time.Minute)
	}
}

func autoBackup() {
	for true {
		fmt.Println("me :备份开始")
		fmt.Println("save hold")
		fn := "backup" + time.Now().Format("06-01-02-15-04")
		CopyDirectory("worlds/Bedrock level", "worlds/"+fn)
		time.Sleep(20 * time.Second)
		fmt.Println("save resume")
		fmt.Println("me :备份结束")
		time.Sleep(24 * time.Hour)
	}
}

func main() {
	time.Sleep(5 * time.Second)
	go printLalala()
	go autoBackup()
	//var tmp string
	cin := bufio.NewReader(os.Stdin)
	for true {
		tmp, err := cin.ReadString('\n')
		if err == nil {
			fmt.Print(tmp)
		}
	}
}
