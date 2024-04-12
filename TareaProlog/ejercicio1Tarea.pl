
es_subconjunto([], _).
es_subconjunto([X|Xs], L) :- member(X, L), es_subconjunto(Xs, L).

subconjunto(Subconjunto, Conjunto) :- es_subconjunto(Subconjunto, Conjunto).

%subconjunto([],[1,2,3,4,5]).
%subconjunto([1,2,3],[1,2,3,4,5]).
%subconjunto([1,2,6],[1,2,3,4,5]).

