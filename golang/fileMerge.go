package main

//import "os"
//
//n 小文件合并大文件
//
//1 资源占用小 性能尽可能好
//2 合并后的文件行与行可以乱序，单行不能错乱
//3 合并完后程序可以退出
//文件名列表已知  io错误忽略 库函数可用
//
//func mergeFiles(fileList []string) outFile string{
//var ch chan
//
//output,_:=os.open("outFileName")
//for i:=0;i<len(fileList);i++{
//go readFile(fileList[i],ch)
//ch <-1
//}
//
//
//
//}
//func readFile(fileName string,ch chan){
//	<-ch
//	input,_ := os.open(fileName)
//	inputData := readall(input) //TODO
//
//	os.write(inputData, output)
//}
