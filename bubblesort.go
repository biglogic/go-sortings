package main 

import "fmt" 

func Bubblesort(list []int) {
	for j := 1 ; j < len(list)  ; j++{
		for x,i := range(list){
			x += 1
			if i > list[j] {
				list[x-1] = list[j]
				list[j] = i  
			}
		} 
	 
 
	}
	fmt.Println(list) 
    
}


func main() {
	a := []int{200,180,430,234,335,640,453,500,345,442,521}
	Bubblesort(a)
}