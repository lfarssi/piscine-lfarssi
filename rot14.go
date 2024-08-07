package piscine

func Rot14(s string) string {
  var r string
   for _, i:= range s {
     if i>= 'a' && i <= 'z' {
       if i>= 'm' {
         r += string(i -  12)
       } else {
         r += string(i + 14)
       }
     } else if i>= 'A' && i<='Z'  {
       if i>= 'M' {
         r += string(i - 12)
       } else {
         r += string(i + 14)
       }
     } else {
       r += string(i)
     }
   }
  return r
}
