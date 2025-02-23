package main
import "fmt"

func main(){
    var nama string
    var target, pendepatan, sisa,hasil float64
    fmt.Scan(&nama, &target, &pendepatan, &sisa)
    target = (100-target)/100
    hasil = pendepatan*target
   
    if hasil == sisa {
        fmt.Printf("%s, kamu berhasil menabung sesuai target, %.2f%% dari pendapatan kamu.\n", nama, target)
    } else if hasil < sisa {
        fmt.Printf("%s, kamu berhasil menabung melebihi target, %.2f%% dari pendapatan kamu.\n", nama, target)
    } else {
        fmt.Printf("%s, kamu gagal menabung sesuai target, %.2f%% dari pendapatan kamu.\n", nama, target)
    }
}
