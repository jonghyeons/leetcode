func average(salary []int) float64 {
    sort.Ints(salary)  
    salary = salary[1:len(salary)-1]
    
    var res float64
    for _, v := range salary {
        res += float64(v)
    }

    return res/float64(len(salary))
}