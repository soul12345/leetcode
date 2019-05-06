package Sudo
func getK(a,b int)int{
    res := (a/3)*3 + b/3
    return res
}

func solveSuduku(board [][]byte, n, m, x, y int, rows,cols,bird []int)bool{
    iX := x
    iY := y
    if y == m{
        iY = 0
        iX = iX + 1
    }
    iK := getK(iX,iY)
    if iX == n{
        for i:=0;i < 9;i++{
            for j:=0; j < 9;j++{
                fmt.Printf("%s ",string(board[i][j]))
            }
            fmt.Printf("\n")
        }
        return true
    }
    if(board[iX][iY] != '.'){
        return solveSuduku(board,n,m,iX,iY+1,rows,cols,bird)
    }
    var val uint
    for val = 1; val <= 9;val++{
        if ((rows[iX]>>val)&1) == 1{
            continue
        }
        if ((cols[iY]>>val)&1) == 1{
            continue
        }
        if ((bird[iK]>>val)&1) == 1{
            continue
        }
        rows[iX] |= (1 << val)
        cols[iY] |= (1 << val)
        bird[iK] |= (1 << val)
        board[iX][iY] = byte(val + '0')
        if(!solveSuduku(board,n,m,iX,iY+1,rows,cols,bird)){
            rows[iX] -= (1<<val)
            cols[iY] -= (1<<val)
            bird[iK] -= (1<<val)
            board[iX][iY] = byte('.')
        }else{
            return true
        }
    }
    return false
}

func solveSudoku(board [][]byte)  {
    n := len(board)
    if n == 0{
        return 
    }
    m := len(board[0])
    rows := make([]int , 9)
    cols := make([]int , 9)
    bird := make([]int , 9)
    for i := 0;i < 9;i++{
        for j := 0; j < 9;j++{
            var val uint = 0
            if board[i][j] != '.'{
                val = uint(board[i][j] - '0')
            }
            rows[i] |= (1 << val)
            cols[j] |= (1 << val)
            k := getK(i,j)
            bird[k] |= (1 << val)
        }
    }
    solveSuduku(board, n, m ,0, 0,rows,cols,bird)
}
