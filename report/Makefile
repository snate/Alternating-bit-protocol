default: relazione.pdf

relazione.pdf: relazione.tex
	pdflatex -shell-escape relazione.tex
	pdflatex -shell-escape relazione.tex

clean:
	rm -f *.lot
	rm -f *.lof
	rm -f *.aux
	rm -f *.toc
	rm -f *.out
	rm -f *.log
	rm -f *.pdf
	rm -f images/*.pdf
	rm -f images/lm/*.pdf
	rm -f images/non-linear/*.pdf
	rm -f images/non-linear/gam/*.pdf
	rm -f images/class/*.pdf
	rm -f sections/*.aux
	