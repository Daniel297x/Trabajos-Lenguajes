(define (money cap i n)
  (cond
    ((= n 0) (display cap))
    ((not(= n 0)) (money (+ cap (* cap i)) i (- n 1)) )))

;pruebas

(money 2000 0.10 0)
(newline)
(money 2000 0.10 1)
(newline)
(money 2000 0.10 2)
(newline)
(money 2000 0.10 3)