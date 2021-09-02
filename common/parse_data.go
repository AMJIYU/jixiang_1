package common
//对字符与数据进行各种处理与转换
import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//删除重复string项目
func RemoveDuplicateString(old []string) []string {
	result := make([]string, 0, len(old))
	temp := map[string]struct{}{}
	for _, item := range old {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//删除重复int
func RemoveDuplicateInt(old []int) []int {
	result := make([]int, 0, len(old))
	temp := map[int]struct{}{}
	for _, item := range old {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//判断slice中是否含有特定int数据
func IntSliceContains(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}


//生成指定长度随机数，只包含字符
func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//切片转换成字符
func SliceToString(slices []string) (result string) {
	b, err := json.Marshal(slices)
	if err != nil {
		return
	}
	result = string(b)
	return
}
// 字符转换成int型切片
func StringToSliceInt(s string) ([]int, error) {
	var r []int
	if s == "" {
		return r, nil
	}
	for _, v := range strings.Split(s, ",") {
		vTrim := strings.TrimSpace(v)
		if i, err := strconv.Atoi(vTrim); err == nil {
			r = append(r, i)
		} else {
			return r, err
		}
	}

	return r, nil
}

// SplitByCharAndTrimSpace 以一个指定字符串进行分割并且移除空格
func SplitByCharAndTrimSpace(s, splitchar string) (result []string) {
	for _, token := range strings.Split(s, splitchar) {
		result = append(result, strings.TrimSpace(token))
	}
	return
}
//讲map所有的key提取变成slice
func ToSlice(m map[string]struct{}) (s []string) {
	for k := range m {
		s = append(s, k)
	}

	return
}




//获取md5 hash值
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

