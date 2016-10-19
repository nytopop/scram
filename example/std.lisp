(def factorial 
	 (lambda (n)
	   (if (eq n 0) 1 
		 (* n (factorial (- n 1))))))
