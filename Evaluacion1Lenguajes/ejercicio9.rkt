(define (eliminar_elemento e lista)
  (apply append (map (lambda (x) (if(not(equal? x e)) (list x) '() )) lista))
  
  ) 

;pruebas
(eliminar_elemento "b" '("f" "a" "d" 1 "c" "b"))
(eliminar_elemento 1 '("f" "a" "d" 1 "c" "b"))
(eliminar_elemento "m" '("f" "a" "d" 1 "c" "b"))
(eliminar_elemento "f" '("f" "a" "d" 1 "c" "b"))