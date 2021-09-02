////时间转换相关函数方法
//
package utils

//
//import (
//	"fmt"
//	"time"
//)
//
//const LayoutFull = "2006-01-02 15:04:05"
//const LayoutYmd = "2006-01-02"
//const LayoutYm = "2006/01"
//const LayoutYm1 = "2006-01"
//
////获取当前时间(返回【2017-04-11 13:24:04】日期格式)
//func CurrentDateTime() string {
//	return time.Now().Format(LayoutFull)
//}
//
////获取当前时间戳
//func CurrentUnixTime() int64 {
//	return time.Now().Unix()
//}
//
////字符串->时间对象
//func Str2Time(formatTimeStr string) time.Time {
//	loc, _ := time.LoadLocation("Local")
//	theTime, _ := time.ParseInLocation(LayoutFull, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型
//	return theTime
//
//}
//
////字符串->时间戳
//func Str2Stamp(formatTimeStr string) int64 {
//	timeStruct := Str2Time(formatTimeStr)
//	millisecond := timeStruct.UnixNano() / 1e6
//	//1592445600
//	return millisecond / 1000
//}
//
////时间对象->字符串
//func Time2Str(t *time.Time) string {
//	return t.Format(LayoutFull)
//}
//
////时间对象->字符串
//func Time2Stamp() int64 {
//	t := time.Now()
//	millisecond := t.UnixNano() / 1e6
//	return millisecond
//}
//
////时间戳->字符串
//func Stamp2Str(stamp int64) string {
//	if stamp == 0 {
//		return ""
//	}
//	return time.Unix(stamp, 0).Format(LayoutFull)
//}
//
////时间戳->UTC
//func Stamp2UTC(stamp int64) string {
//	if stamp == 0 {
//		return ""
//	}
//
//	return Stamp2Time(stamp).UTC().Format("2006-01-02T15:04:05Z")
//}
//
////时间戳转年月日
//func DateFormatYmd(timestamp int) string {
//	tm := time.Unix(int64(timestamp), 0)
//	return tm.Format(LayoutYmd)
//}
//
////当前年月日
//func CurrentDate() string {
//	tm := time.Unix(time.Now().Unix(), 0)
//	return tm.Format(LayoutYmd)
//}
//
////时间戳->时间对象
//func Stamp2Time(stamp int64) time.Time {
//	stampStr := Stamp2Str(stamp)
//	timer := Str2Time(stampStr)
//	return timer
//}
//
////获取零点时间戳
//func DateZeroUnix(addDays int) int64 {
//	zero, _ := time.ParseInLocation(LayoutYmd, time.Now().Format(LayoutYmd), time.Local)
//	if addDays == 0 { //当天零点时间戳
//		return zero.Unix()
//	}
//
//	return zero.AddDate(0, 0, addDays).Unix() //其他时间，可加可减
//}
//
////查询指定时间的日期
//func DateSpecifyTime(day int) string {
//	nTime := time.Now()
//	yesTime := nTime.AddDate(0, 0, day)
//	return yesTime.Format(LayoutYmd)
//}
//
////SecondsFormatHHMMSS 秒数转时分秒，,如果HH 为0，则删掉
//func SecondsFormatHHMMSS(seconds int64) string {
//	HH := "00"
//	MM := "00"
//	SS := "00"
//	h := seconds / 3600
//	if h < 10 {
//		HH = "0" + ConvertToString(h)
//	} else {
//		HH = ConvertToString(h)
//	}
//
//	m := seconds / 60 % 60
//	if m < 10 {
//		MM = "0" + ConvertToString(m)
//	} else {
//		MM = ConvertToString(m)
//	}
//
//	s := seconds % 60
//	if s < 10 {
//		SS = "0" + ConvertToString(s)
//	} else {
//		SS = ConvertToString(s)
//	}
//	if HH == "00" {
//		return MM + ":" + SS
//	}
//	return HH + ":" + MM + ":" + SS
//}
//
//// SecondsFormatHHMMSS 秒数转时分秒 ,如果HH 为0，则不转换
//func SecondsFormatHHMMSSNotTrans(seconds int64) string {
//	HH := "00"
//	MM := "00"
//	SS := "00"
//	h := seconds / 3600
//	if h < 10 {
//		HH = "0" + ConvertToString(h)
//	} else {
//		HH = ConvertToString(h)
//	}
//
//	m := seconds / 60 % 60
//	if m < 10 {
//		MM = "0" + ConvertToString(m)
//	} else {
//		MM = ConvertToString(m)
//	}
//
//	s := seconds % 60
//	if s < 10 {
//		SS = "0" + ConvertToString(s)
//	} else {
//		SS = ConvertToString(s)
//	}
//	return HH + ":" + MM + ":" + SS
//}
//
////TimeInfoFromStamp 根据时间戳得到详细的时间信息（年，月，日，时，分，秒）
//func TimeInfoFromStamp(stamp int64) (int, int, int, int, int, int) {
//	dstTime := Stamp2Time(stamp)
//	return dstTime.Local().Year(), int(dstTime.Local().Month()), dstTime.Local().Day(), dstTime.Local().Hour(), dstTime.Local().Minute(), dstTime.Local().Second()
//}
//
////StrTimeInfoFromStamp 根据时间戳得到详细的时间信息（年，月，日，时，分，秒）
//func StrTimeInfoFromStamp(stamp int64) (string, string, string, string, string, string) {
//	dstTime := Stamp2Time(stamp)
//
//	szHour := fmt.Sprintf("%d", dstTime.Local().Hour())
//	szMinute := fmt.Sprintf("%d", dstTime.Local().Minute())
//	if len(fmt.Sprintf("%d", dstTime.Local().Hour())) < 2 {
//		szHour = fmt.Sprintf("0%d", dstTime.Local().Hour())
//	}
//	if len(fmt.Sprintf("%d", dstTime.Local().Minute())) < 2 {
//		szMinute = fmt.Sprintf("0%d", dstTime.Local().Minute())
//	}
//	return fmt.Sprintf("%d", dstTime.Local().Year()),
//		fmt.Sprintf("%d", int(dstTime.Local().Month())),
//		fmt.Sprintf("%d", dstTime.Local().Day()),
//		szHour,
//		szMinute,
//		fmt.Sprintf("%d", dstTime.Local().Second())
//}
//
////获取设置天数到当前时间之差
//func TwoTimeDifference(day int) time.Duration {
//
//	if day < 0 {
//		return 0
//	}
//	lastTimeUnix := DateZeroUnix(day)
//	nowTimeUnix := time.Now().Unix()
//	return time.Duration(lastTimeUnix - nowTimeUnix)
//}
