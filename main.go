package main

import(
    "fmt"
    "net/http"

)

func main(){
    // "/"��η� �������� �� ó���� �ڵ鷯 �Լ� ����
    http.HandleFunc("/"
                    , func(w http.ResposeWriter, r *http.Request){
                    //"welcome!" ���ڿ��� ȭ�鿡 ���
                    fmt.Fprintln(W, "welcome")
                    }
    )
    
    // 8080 ��Ʈ�� �� ���� ����
    http.ListenAndServe(":8080", nil)
    
}