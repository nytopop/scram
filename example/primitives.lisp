;; PRIMITIVES
; quote
(quote hello)
(quote (hello, world!))
(quote (a b c d))

; if
(if (eq 1 1) (quote conseq) (quote alt))
(if (eq 0 1) (quote conseq) (quote alt))

; cond
(cond 
  ((eq 0 1) (quote nope))
  ((eq 1 1) (quote yep)))

; def
(def x 44)

; set
(set x 33)

; lambda
(def add3
	 (lambda (a b c) (+ a b c)))
(add3 11 22 33)

; begin
(begin add3)
(begin x)

; go / chan / -> / <-
(def concurrent-add3 
	 (lambda (ch a b c) (-> ch (+ a b c))))
(def addchan (chan 1))
(go (concurrent-add3 addchan 192 124 234))
(<- addchan)

;; BUILTINS
; cons
(cons 1 (cons 2 (cons 3 4)))

; car
(car (list 1 2 3 4))

; cdr
(cdr (list 1 2 3 4))

; list
(list 1 2 3 4)

; +
(+ 1 2 3)
(+ 1 2 3 4 5 6 7 8 9)

; -
(- 1 2 3)
(- 1 2 3 4 5 6 7 8 9)

; *
(* 1 2 3)
(* 1 2 3 4 5 6 7 8 9)

; /
(/ 1 2 3)
(/ 1 2 3 4 5 6 7 8 9)

; eq
(eq 1 0)
(eq 1 1)
(eq (list 1 2 3) (list 1 2))
(eq (list 1 2) (list 1 2))

; lt
(lt 1 2)
(lt 2 2)

; lte
(lte 1 2)
(lte 2 2)

; gt
(gt 2 1)
(gt 2 2)

; gte
(gte 2 1)
(gte 2 2)
