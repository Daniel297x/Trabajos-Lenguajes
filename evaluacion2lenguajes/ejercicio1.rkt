(define
  ListaProductos
  '(("arroz" 8 1800) ("frijoles" 5 1200) ("azucar" 6 1100) ("cafe" 2 2800) ("leche" 9 1100) ))

(define (agregarProducto Lista nombre cantidad precio)
  (cond ((null? Lista)
         (list (list nombre cantidad precio)))
        ((equal? nombre (caar Lista))
         (cons (list (caar Lista) (+ (cadar Lista) cantidad ) precio) (cdr Lista)))
        (else
         ( cons (car Lista)  (agregarProducto (cdr Lista) nombre cantidad precio)))))

(define (venderProducto Lista nombre cantidad )
  (cond ((null? Lista) (display "no existe el producto") '())
        ((equal? nombre (caar Lista)) (cons (list (caar Lista) (- (list-ref (car Lista) 1) cantidad)  (list-ref (car Lista) 2)) (cdr Lista)))
  (else (cons (car Lista) (venderProducto (cdr Lista) nombre cantidad)))))

(define (existenciaMinimas Lista cantidad)
  (filter (lambda (x) (<= (cadr x) cantidad)) Lista))

(define (crearFactura listaP listaV)
  (cond
    ((null? listaP)  '()  )
    
    ((not(equal? (car listaV) (car listaP) ))  (cons (list (list-ref (car listaV) 0) (- (list-ref (car listaP) 1) (list-ref (car listaV) 1)) (list-ref (car listaV) 2)) (crearFactura (cdr listaP) (cdr listaV))))
    (else 
     (crearFactura (cdr listaP) (cdr listaV)))) 
    )
(define (calcularImpuesto listaFac montoI t)
  (cond
    ((null? listaFac) (display "total impuestos: ") (display t))
    ((>= (list-ref (car listaFac) 2) montoI ) (calcularImpuesto (cdr listaFac) montoI (+ t (* (list-ref (car listaFac) 1) (* (list-ref (car listaFac) 2) 0.13))) ))
    (else
     (calcularImpuesto (cdr listaFac) montoI(+ t 0) ))
    ))

(define (calcularTotal listaFac t)
  (cond
    ((null? listaFac) (display "total: ") (display t))
    
    (else
     (calcularTotal (cdr listaFac) (+ t (* (list-ref (car listaFac) 1) (list-ref (car listaFac) 2)) ) ))
    ))




;pruebas
(display "Factura 1:")
(crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "frijoles" 2) "arroz" 2))
(calcularImpuesto (crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "frijoles" 2) "arroz" 2)) 1200 0)
(newline )
(calcularTotal (crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "frijoles" 2) "arroz" 2)) 0)
(newline )(newline )
(display "Factura 2:")
(crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "leche" 5) "arroz" 2))
(calcularImpuesto (crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "leche" 5) "arroz" 2)) 1200 0)
(newline )
(calcularTotal (crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "leche" 5) "arroz" 2)) 0)
(newline )(newline )
(display "Factura 3:")
(crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "frijoles" 2) "cafe" 1))
(calcularImpuesto (crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "frijoles" 2) "cafe" 1)) 1200 0)
(newline )
(calcularTotal (crearFactura ListaProductos (venderProducto (venderProducto ListaProductos "frijoles" 2) "cafe" 1)) 0)