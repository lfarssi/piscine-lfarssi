package piscine

func Rot14(s string) string {
  var r string
   for _, i:= range s {
     if (i>= 'a' && i <= 'z') || (i>= 'A' && i<='Z') {
       r = i + 14
     }
   }
  return r
}
