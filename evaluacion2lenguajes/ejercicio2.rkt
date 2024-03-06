(define (separar texto)
  (define (split-helper str acc)
    (if (string=? str "") 
        (reverse acc)
        (let loop ((remaining str)
                   (word "")
                   (words acc))
          (cond
            ((string=? remaining "") 
             (reverse (cons word words)))
            ((char=? (string-ref remaining 0) #\space) 
             (loop (substring remaining 1 (string-length remaining)) "" (cons word words)))  
            (else 
             (loop (substring remaining 1 (string-length remaining)) 
                   (string-append word (string (string-ref remaining 0))) 
                   words))))))

  (split-helper texto '()))


(define(validar subc palabras)
  (cond
    ((null? palabras) #f)
    ((equal? (car palabras) subc) #t)
    (else (validar subc (cdr palabras))))
  )

(define (subCadena subc lista)
  (filter (lambda (x) (validar subc (separar x)) ) lista))

;prueba
(subCadena "la" '("la casa" "el perro" "pintando la cerca"))
