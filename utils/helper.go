package utils

//
////获取
//func ArrayColumn(input map[string]map[string]interface{}, key string) (res []interface{}) {
//	res = make([]interface{}, 0, len(input))
//	for _, val := range input {
//		if v, ok := val[key]; ok {
//			res = append(res, v)
//		}
//	}
//	return
//}
//
///*
// * Gzip压缩处理
// */
//func GzipEncode(in []byte) ([]byte, error) {
//	var (
//		buffer bytes.Buffer
//		out    []byte
//		err    error
//	)
//	//writer := gzip.NewWriter(&buffer)
//	//使用压缩等级
//	writer, _ := gzip.NewWriterLevel(&buffer, 1)
//
//	//最开始一直把关闭放在了defer里面，但是这样压缩后数据为空
//	//查阅相关资料才知道：gzip压缩的过程中，Write之后一定要及时Close，不能defer，这样才能flush，否则得不到任何数据
//	/*defer func() {
//		err := writer.Close()
//		if err != nil {
//			logrus.Info("Gzip压缩失败", err.Error())
//		}
//	}()*/
//	_, err = writer.Write(in)
//	if err != nil {
//		err = writer.Close()
//		if err != nil {
//			logrus.Info("Gzip压缩关闭失败", err.Error())
//		}
//		return out, err
//	}
//	if err = writer.Close(); err != nil {
//		logrus.Info("Gzip压缩关闭失败", err.Error())
//		return out, err
//	}
//	return buffer.Bytes(), nil
//}
//
///*
// * Gzip解压处理
// */
//func GzipDecode(in []byte) ([]byte, error) {
//	reader, err := gzip.NewReader(bytes.NewReader(in))
//	if err != nil {
//		var out []byte
//		return out, err
//	}
//	defer func() {
//		if err = reader.Close(); err != nil {
//			logrus.Info("Gzip解压关闭失败", err.Error())
//		}
//	}()
//
//	return ioutil.ReadAll(reader)
//}
//
////检查请求是否重复(防止暴力点击)
//func CheckRepeatRequest(key string, expireTime int64) bool {
//	// 缓存的key
//	cacheKey := "cache_repeat_request:" + key
//
//	if expireTime < 1 {
//		dao.Redis.Del(cacheKey)
//		return true
//	}
//
//	if dao.Redis.Incr(cacheKey).Val() == 1 {
//		dao.Redis.Expire(cacheKey, time.Duration(expireTime)*time.Second)
//		return false
//	}
//
//	return true
//}
//
////删除重复请求
//func DelCheckRepeatRequest(key string) {
//	cacheKey := "cache_repeat_request:" + key
//	dao.Redis.Del(cacheKey)
//}
//
//// 转换为字符串类型
//func ConvertToString(i interface{}) string {
//	switch v := i.(type) {
//	case string:
//		return v
//	case int:
//		return strconv.Itoa(v)
//	case int64:
//		return strconv.FormatInt(v, 10)
//	case uint64:
//		return strconv.FormatUint(v, 10)
//	case float32:
//		return strconv.FormatFloat(float64(v), 'f', -1, 32)
//	case float64:
//		return strconv.FormatFloat(v, 'f', 0, 64)
//	case bool:
//		if i == true {
//			return "1"
//		}
//		return ""
//	default:
//		logrus.Error("处理请求参数时出现了未知数据类型", "=====", i, "====", v)
//	}
//
//	return ""
//}
//
////str转uint64
//func Str2UInt64(v string) uint64 {
//	if v == "" {
//		return 0
//	}
//	i, err := strconv.ParseUint(v, 10, 64)
//	if err != nil {
//		logrus.Error(err.Error(), "参数值为：", v)
//	}
//
//	return i
//}
//
//// string转int64
//func String2Int64(v string) int64 {
//	if v == "" {
//		return 0
//	}
//	i, err := strconv.ParseInt(v, 10, 64)
//	if err != nil {
//		logrus.Error(err.Error(), "参数值为：", v)
//	}
//
//	return i
//}
//
//// 生成min~max的随机数
//func GenerateRangeNum(min, max int64) int64 {
//	//如果最大值与最小值相等，就直接返回
//	if min == max {
//		return min
//	}
//
//	rand.Seed(time.Now().Unix())
//	randNum := rand.Int63n(max+1-min) + min //+1是因为Int63n随机出来是不包含最大值
//	return randNum
//}
//
//// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
//const (
//	XForwardedFor = "X-Forwarded-For"
//	XRealIP       = "X-Real-IP"
//)
//
//func RemoteIp(req *http.Request) string {
//	remoteAddr := req.RemoteAddr
//	if ip := req.Header.Get(XRealIP); ip != "" {
//		remoteAddr = ip
//	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
//		remoteAddr = ip
//	} else {
//		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
//	}
//
//	if remoteAddr == "::1" {
//		remoteAddr = "127.0.0.1"
//	}
//
//	return remoteAddr
//}
//
////三元表达式处理
//func Ternary(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
//	if condition == true {
//		return trueValue
//	}
//	return falseValue
//}
//
////字符串转Unicode ,emoji表情转utf16
////outputType 编码方式  json  html
//func StringToUnicode(text string, outputType string) (data string) {
//	rs := []rune(text)
//	for _, r := range rs {
//		st := string(r)
//		if len(st) == 4 {
//			//处理emoji表情
//			ss := utf16.Encode([]rune(st))
//			for _, r := range ss {
//				data += "\\u" + strconv.FormatInt(int64(r), 16)
//			}
//			continue
//		}
//		rint := int(r)
//		if rint < 128 {
//			if outputType == "json" {
//				data += st
//			}
//			if outputType == "html" {
//				data += st
//			}
//		} else {
//			if outputType == "json" {
//				//data += "\\u" + strconv.FormatInt(int64(rint), 16) // json
//				textQuoted := strconv.QuoteToASCII(st)
//				textUnquoted := textQuoted[1 : len(textQuoted)-1]
//				data += textUnquoted
//			}
//			if outputType == "html" {
//				data += "&#" + strconv.Itoa(int(r)) + ";" // 网页
//			}
//		}
//	}
//	return
//}
//
////md5加密函数
//func Md5(str string) string {
//	data := []byte(str)
//	has := md5.Sum(data)
//	md5str := fmt.Sprintf("%x", has)
//	return md5str
//}
//
////正则表达式提取一个字符串中的字母/数字/_组成一个字符串.有长度限制。使用时请注意
//func RegexpGetString(s string) string {
//	regexpString := "[A-Za-z0-9_]+"
//	regexpM, err := regexp.Compile(regexpString)
//	if err != nil {
//		logrus.Error(err)
//		return ""
//	}
//
//	findByteSlice := regexpM.FindAll([]byte(s), 30)
//	if len(findByteSlice) == 0 {
//		return ""
//	}
//	data := ""
//	for _, findByte := range findByteSlice {
//		data += string(findByte)
//	}
//	return data
//}
//
////截取一个字符串的长度
//// str:字符串
//// length:前几个字
//func SubStrByByte(str string, length int) string {
//	rune := []rune(str)
//	if length >= len(rune) {
//		length = len(rune)
//	}
//	str = string(rune[0:length])
//	return str
//}
//
////获取本机的一个网络mac地址
//func GetMac() string {
//	interfaces, err := net.Interfaces()
//	if err != nil {
//		logrus.Error("获取网络mac地址获得一个错误:" + err.Error())
//		return ""
//	}
//	for _, inter := range interfaces {
//		mac := inter.HardwareAddr.String()
//		if mac != "" {
//			return mac
//		}
//	}
//	logrus.Error("未能正确获取本机mac地址!")
//	return ""
//}
//
////把秒格式为"时:分:秒"
//func FormatTimeFromSeconds(format string, seconds int64) string {
//
//	second := fmt.Sprintf("%.2d", seconds%60)
//
//	formatTime := ""
//	if format == "H:i:s" {
//		hour := fmt.Sprintf("%.2d", seconds/3600)
//		minute := fmt.Sprintf("%.2d", (seconds%3600)/60)
//		formatTime = fmt.Sprintf("%v:%v:%v", hour, minute, second)
//		return formatTime
//	}
//
//	if format == "i:s" {
//		minute := fmt.Sprintf("%.2d", seconds/60)
//		formatTime = fmt.Sprintf("%v:%v", minute, second)
//	}
//
//	if format == "mm:ss" {
//		minute := fmt.Sprintf("%.2d", seconds/60)
//		formatTime = fmt.Sprintf("%v分%v秒", minute, second)
//	}
//
//	return formatTime
//}
//
////版本号比较
//func VersionCompare(v1, v2, operator string) bool {
//	com := compare(v1, v2)
//	switch operator {
//	case "==":
//		if com == 0 {
//			return true
//		}
//	case "<":
//		if com == 2 {
//			return true
//		}
//	case ">":
//		if com == 1 {
//			return true
//		}
//	case "<=":
//		if com == 0 || com == 2 {
//			return true
//		}
//	case ">=":
//		if com == 0 || com == 1 {
//			return true
//		}
//	}
//	return false
//}
//
//func compare(v1, v2 string) int {
//	// 替换一些常见的版本符号
//	replaceMap := map[string]string{"V": "", "v": "", "-": "."}
//	//keywords := {"alpha,beta,rc,p"}
//	for k, v := range replaceMap {
//		if strings.Contains(v1, k) {
//			strings.Replace(v1, k, v, -1)
//		}
//		if strings.Contains(v2, k) {
//			strings.Replace(v2, k, v, -1)
//		}
//	}
//	ver1 := strings.Split(v1, ".")
//	ver2 := strings.Split(v2, ".")
//	// 找出v1和v2哪一个最短
//	var shorter int
//	if len(ver1) > len(ver2) {
//		shorter = len(ver2)
//	} else {
//		shorter = len(ver1)
//	}
//	// 循环比较
//	for i := 0; i < shorter; i++ {
//		if ver1[i] == ver2[i] {
//			if shorter-1 == i {
//				if len(ver1) == len(ver2) {
//					return 0
//				} else {
//					if len(ver1) > len(ver2) {
//						return 1
//					} else {
//						return 2
//					}
//				}
//			}
//		} else if ver1[i] > ver2[i] {
//			return 1
//		} else {
//			return 2
//		}
//	}
//	return -1
//}
//
////限制次数
//func RateLimiter(cacheKey string, number int64, expireTime int64) (bool, error) {
//	if cacheKey == "" {
//		return false, errors.New("缓存Key不能为空")
//	}
//	if number <= 0 {
//		return false, errors.New("次数不能小于0")
//	}
//	if expireTime <= 0 {
//		return false, errors.New("缓存时间不能小于0")
//	}
//	num := dao.Redis.Incr(cacheKey).Val()
//	if num == 1 {
//		dao.Redis.Expire(cacheKey, time.Duration(expireTime)*time.Second)
//	}
//	if num > number {
//		return true, nil
//	}
//	return false, nil
//}
//
////格式化时间戳
//func FormatTimeStampToStr(timestamp int64, format string) string {
//	t := time.Unix(timestamp, 0)
//	formats := map[string]string{
//		"%Y": "2006",
//		"%m": "01",
//		"%d": "02",
//		"%H": "15",
//		"%i": "04",
//		"%s": "05",
//	}
//
//	for k, v := range formats {
//		format = strings.Replace(format, k, v, -1)
//	}
//	return t.Format(format)
//}
//
//// 生成随机一个长度的字符串
//func RandString(size int) string {
//	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
//	var s bytes.Buffer
//	for i := 0; i < size; i++ {
//		s.WriteByte(char[rand.Int63()%int64(len(char))])
//	}
//	return s.String()
//}
//
//// 返回2个切片的交集
//func SliceIntersect(s1, s2 []string) (s []string) {
//	for _, v1 := range s1 {
//		canAdd := false
//		for _, v2 := range s2 {
//			if v1 == v2 {
//				canAdd = true
//			}
//		}
//		if canAdd {
//			s = append(s, v1)
//		}
//	}
//
//	return s
//}
//
//// 发送POST请求
//// url：         请求地址
//// data：        POST请求提交的数据
//// contentType： 请求体格式，如：application/json
//// content：     请求放回的内容
//func HttpPost(url string, data interface{}, contentType string) string {
//	// 超时时间：600秒 预加载耗时过长
//	client := &http.Client{Timeout: 600 * time.Second}
//	jsonStr, _ := json.Marshal(data)
//	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		logrus.Error(err)
//		return ""
//	}
//	defer resp.Body.Close()
//
//	result, _ := ioutil.ReadAll(resp.Body)
//	return string(result)
//}
//
//// 发送GET请求
//// url：         请求地址
//// data：        POST请求提交的数据
//// content：     请求放回的内容 ?test=1&a=1
//func HttpGet(url string, path string) string {
//	// 超时时间：5秒
//	client := &http.Client{Timeout: 5 * time.Second}
//	uri := url + path
//	resp, err := client.Get(uri)
//	if err != nil {
//		logrus.Error(err)
//		return ""
//	}
//	defer resp.Body.Close()
//
//	result, _ := ioutil.ReadAll(resp.Body)
//	return string(result)
//}
//
//// Implode implode()
//func Implode(glue string, pieces []string) string {
//	var buf bytes.Buffer
//	l := len(pieces)
//	for _, str := range pieces {
//		buf.WriteString(str)
//		if l--; l > 0 {
//			buf.WriteString(glue)
//		}
//	}
//	return buf.String()
//}
//
////
//func GetInteractType(EventType int) string {
//	switch EventType {
//	case 1:
//		return "h5"
//	case 2:
//		return "answer_sheet"
//	case 3:
//		return "classroom_test"
//	case 4:
//		return "photo_to_wall"
//	default:
//		return ""
//	}
//}
//
////转换数组int64转数组字符串
//func ArrInt64ArrString(input []int64) []string {
//	var arrSting []string
//	for _, value := range input {
//		arrSting = append(arrSting, strconv.FormatInt(value, 10))
//	}
//	return arrSting
//}
//
////转换数组字符串转数组int64
//func ArrStringArrInt64(input []string) []int64 {
//	var arrSting []int64
//	for _, value := range input {
//		num, _ := strconv.ParseInt(value, 10, 64)
//		arrSting = append(arrSting, num)
//	}
//	return arrSting
//}
//
////按key进行排序
//func KSort(input map[string]interface{}) (data map[string]interface{}) {
//	data = map[string]interface{}{}
//	var keys []string
//	for key, _ := range input {
//		keys = append(keys, key)
//	}
//	sort.Strings(keys)
//	for _, key := range keys {
//		data[key] = input[key]
//	}
//	return
//}
//
////转换成：test=1&test2=2
//func HttpQueryBuild(input map[string]string) string {
//	var httpParameters []string
//	for key, value := range input {
//		httpParameters = append(httpParameters, key+"="+value)
//	}
//	return strings.Join(httpParameters, "&")
//}
//
////分页数据
//func Pagination(data interface{}, total int64, page, perPage int) gin.H {
//	return gin.H{
//		"data": data,
//		"meta": gin.H{
//			"pagination": gin.H{
//				"total":     total,
//				"count":     total,
//				"per_page":  perPage,
//				"page":      page,
//				"last_page": int64(math.Ceil(float64(total) / float64(perPage))),
//			},
//		},
//	}
//}
//
////获取应用Url
//func GetAppUrl() string {
//	return viper.GetString("app.http_scheme") + "://" + viper.GetString("app.host")
//}
//
////获取应用Url
//func GetInsideAppName() string {
//	return "http://" + viper.GetString("app.xthk_app_name")
//}
//
////获取wsUrl
//func GetWsUrl(uri string) (url string) {
//	url = viper.GetString("ws.scheme") + "://" + viper.GetString("ws.host")
//	if viper.GetString("app.app_env") == "local" {
//		url += fmt.Sprintf(":%d", viper.GetInt("ws.port"))
//	}
//	url += uri
//	return
//}
//
//// 打包成zip文件
//// srcDir 打包的目录
//// zipfileName zip文件名称
//// isDelete  是否删除打包的目录
//func CompressedFile(srcDir string, zipfileName string, isDelete bool) error {
//	var err error
//	//支持的扩展名
//	var extension = []string{
//		"zip",
//	}
//	var str = zipfileName[len(zipfileName)-3:]
//	if arrays.ContainsString(extension, str) == -1 {
//		return errors.New("当前不支持" + str)
//	}
//	// 预防：旧文件无法覆盖
//	err = os.RemoveAll(zipfileName)
//	if err != nil {
//		return err
//	}
//	// 创建：zip文件
//	zipfile, _ := os.Create(zipfileName)
//
//	// 打开：zip文件
//	archive := zip.NewWriter(zipfile)
//
//	// 遍历路径信息
//	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
//		// 如果是源路径，提前进行下一个遍历
//		if path == srcDir {
//			return nil
//		}
//
//		header, _ := zip.FileInfoHeader(info)
//		path = strings.ReplaceAll(path, "\\", "/")     // 对同时兼容Linux 和 win 进行处理
//		header.Name = strings.TrimPrefix(path, srcDir) // +`/`
//		//fmt.Println(header.Name, path, srcDir)         // pictures\10.jpg E:xx\V1\pictures\10.jpg E:\xx\V1
//		// 判断：文件是不是文件夹
//		if info.IsDir() {
//			header.Name += `/`
//		} else {
//			// 设置：zip的文件压缩算法
//			header.Method = zip.Deflate
//		}
//
//		// 创建：压缩包头部信息
//		writer, _ := archive.CreateHeader(header)
//		if !info.IsDir() {
//			file, _ := os.Open(path)
//			defer func() {
//				err = file.Close()
//				if err != nil {
//					logrus.Error("关闭失败", err.Error())
//				}
//			}()
//			_, err := io.Copy(writer, file)
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//	if err != nil {
//		return err
//	}
//	defer func() {
//		//关闭一托依赖
//		//zipfile 创建zip包
//		err = zipfile.Close()
//
//		//打开zip文件
//		err = archive.Close()
//		if err != nil {
//			logrus.Error("关闭失败", err.Error())
//		}
//		//删除打包的目录
//		if isDelete == true {
//			err = os.RemoveAll(srcDir)
//			logrus.Error("删除失败", err.Error())
//		}
//	}()
//	return nil
//}
//
////转map
//func ToStructByMap(input interface{}) (data map[string]interface{}) {
//	inputByByte, _ := json.Marshal(input)
//	_ = json.Unmarshal(inputByByte, &data)
//	return
//}
//
////转string
//func ToInterfaceByString(input interface{}) (data string) {
//	inputByByte, _ := json.Marshal(input)
//	data = string(inputByByte)
//	return
//}
