# Set of minimalistic style packages

### 1 Spinner
How to use:
   ```golang
    fmt.Println("Work in progress")
  	s := spinner.Start()
    time.Sleep(time.Second * 10)
    s.Stop()
   ```
![](images/spinner.gif)