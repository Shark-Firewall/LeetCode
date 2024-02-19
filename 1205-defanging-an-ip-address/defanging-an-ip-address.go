func defangIPaddr(address string) string {
   var ipAddress string = "";
   for _, j := range address {
       if (string(j) == "."){
           ipAddress += "[.]";
       }else{
           ipAddress += string(j);
       }
   } 
   return ipAddress;
}