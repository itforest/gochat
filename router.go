type router struct{
    // Ű : http �޼���
    // �� : URL ���Ϻ��� ������ HandleFunc
    handlers map[string]map[string]http.HandlerFunc
}