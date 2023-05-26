package Array

func Preppend(arr[] int,value int)[] int{
	newarr:=make([] int,len(arr)+1)
	for i := 0; i < len(arr); i++ {
		newarr[i+1]=arr[i]
		
	}
	newarr[0]=value
	return newarr
}