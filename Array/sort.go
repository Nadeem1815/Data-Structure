package Array

func Sort(arr []int)[] int{
	
	for i := 0; i < len(arr)-1; i++ {
		for j := i+1;j < len(arr); j++ {
			if arr[i]>arr[j] {
				temp:=arr[i]
				arr[i]=arr[j]
				arr[j]=temp
			}
			
		}
	}
	return arr
}