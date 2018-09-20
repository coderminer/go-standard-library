### Go中关于文件相关的操作

#### 创建一个空文件

```
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal("create file err", err)
	}
	log.Println(file)
	file.Close()
}
```

#### 获取文件的信息

```
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type:%T\v\n", fileInfo.Sys())
	fmt.Printf("System info:%+v\n\n", fileInfo.Sys())
}
```

#### 重命名和移动文件

```
package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "empty.txt"
	newPath := "test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
```

#### 删除文件

```
package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
}

```

#### 打开关闭文件

```
package main

import (
	"log"
	"os"
)

func main() {
	//简单的打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
```

打开文件的一些其他的操作

```
os.O_RDONLY     //只读
os.O_WRONLY     //只写
os.O_RDWR       //读写文件
os.O_APPEND     //追加文件
os.O_CREATE     //不存在时创建文件
os.O_TRUNC      //打开时截断文件
```

#### 检查文件是否存在

```
package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist.File information:")
	log.Println(fileInfo)
}
```

#### 检查文件的读写权限

```
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error:Write permission denied.")
		}
	}
	file.Close()
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error:Read permission denied.")
		}
	}
	file.Close()
}
```

#### 改变文件的权限、所有权和时间戳

```
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	//改变权限
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	//改变所有权 适用于linux, windows不支持
	err = os.Chown("test.txt", os.Getuid(), os.Getegid())
	if err != nil {
		log.Println(err)
	}

	//改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

```

#### 复制文件

```
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//打开原文件
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	//创建新文件
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//文件复制
	bytes, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytes)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
```

#### 移动位置

```
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var offset int64 = 5

	var whence int = 0
	newPos, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPos)

	newPos, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPos)

	currentPos, err := file.Seek(0, 1)
	fmt.Println("current pos:", currentPos)

	newPos, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("position after seek 0,0:", newPos)
}

```

#### 向文件中写入字节

```
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := []byte("测试写入功能！")
	bw, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bw)
}

```

#### 快速写入文件

```
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("test.text", []byte("测试快速写入功能!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
```

#### 在写入时使用缓存

```
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := bufio.NewWriter(file)
	bw, err := buffer.Write([]byte{65, 66, 67})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written:%d\n", bw)
	bw, err = buffer.WriteString("\n写入字符串")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bw)

	unFlushedBuffer := buffer.Buffered()
	log.Printf("Bytes buffered:%d\n", unFlushedBuffer)

	ba := buffer.Available()
	log.Printf("Available buffer: %d\n", ba)

	buffer.Flush()

	buffer.Reset(buffer)

	ba = buffer.Available()
	log.Printf("Availabled buffer:%d\n", ba)

	buffer = bufio.NewWriterSize(buffer, 8000)
	ba = buffer.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", ba)

}

```

#### 从文件中读取n个字节

```
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
    //从文件中读取16个字节
	bytes := make([]byte, 16)
	br, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of bytes read: %d\n", br)
	log.Printf("Data read: %s\n", bytes)
}

```

#### 读取文件中全部内容

```
data, err := ioutil.ReadAll(file)
```

#### 快速读取文件到内存中

```
package main

import (
    "log"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Data read: %s\n", data)
}
```

[更多精彩内容](http://www.coderminer.com)