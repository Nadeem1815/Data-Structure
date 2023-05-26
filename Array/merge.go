package Array

func Merge(arr[] int,arr1 [] int)[]int{
	newarr:=make([]int,len(arr)+len(arr1))
	for i := 0; i < len(arr); i++ {
		newarr[i]=arr[i]
		
	}
	for i := len(arr); i < len(newarr); i++ {
		newarr[i]=arr1[i-len(arr)]
		
	}
	return newarr
}