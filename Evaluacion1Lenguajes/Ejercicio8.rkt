(define (sub-conjunto? sconj conj)
  
  (equal? sconj (apply append (map (lambda (x) (filter(lambda (n) (equal? x n)) conj)) sconj)))
  )
    
;pruebas
(sub-conjunto? '("a" "b" "c" "f" 1) '("f" "a" "d" 1 "c" "b"))
(sub-conjunto? '("a" "b" "c" "f" 1) '("r" "p" "d" 1 "c" "l"))

