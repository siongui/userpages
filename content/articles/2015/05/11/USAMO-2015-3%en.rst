[Math] USAMO 2015 Problem 3
###########################

:date: 2015-05-11T19:00-08:00
:author: Shen-Fu Tsai
:tags: Math
:category: Math
:summary: USAMO 2015 Problem 3

As high school math competition gets harder and harder, it becomes notable for me to independently
solve solve a 3/6 problem at IMO. Given USAMO appears, in my opinion, to be at least as hard as IMO,
I should write this down.

Problem: Let :math:`n` be a positive integer. We color each subset of :math:`N=\{1,2,...,n\}` with blue or red, including empty set. Given a subset :math:`A` of :math:`N`, let :math:`f(A)` be the number of blue subsets of :math:`A`. Find :math:`T(n)`, the number of colorings such that for any subsets :math:`A` and :math:`B` of :math:`N`.

f(A)f(B)=f(A intersect B)f(A union B)	(*)

At first I wasn't sure whether it's a combinatoric or functional problem, though there is never a need, and actually bad, for every problem to belong to a specific category because that gives people some non-intuitive clue to the approach.  After spending quite some time working on it as a functional problem, I didn't seem to get anywhere.  I have to admit that I then shamelessly wrote a simple code to compute some values of T(n) and then guessed the pattern and prove it, which is a great advantage, and never possible, in real competition.

Anyways, the key observation is this: what will happen if the empty set is red and f({})=0? Then one of f({1}),f({2}),...,f({n}) is zero. Let's say f({1})=0, then one of f({1,2}),f({1,3}),...,f({1, n}) is zero. Let's say f({1,2})=0, then...

Eventually, with f({})=0 we arrive at a situation that without loss of generality, for any subset A of {1,2,...,n-1} its color is red and f(A)=0, and all sets that may have non-zero f() value are those including n. Moreover, interestingly we could treat these 2^(n-1) sets containing n as another identical problem on {1,2,3,...,n-1}, as long as n is temporarily ignored. Now either f({n})=1, or f({n})=0 and maybe some set containing n is not 0, and so on. So if G(n) is the number of coloring with f({})=1, we *might* have

T(n)=C(n,0)G(n)+C(n,1)G(n-1)+C(n,2)G(n-2)+...+C(n,n)G(0)+1

The rough argument is to have a minimum set A with f(A)=1. For |A|=0 the count is C(n,0)G(n), for |A|=1 the count is C(n,1)G(n-1), and so on. 1 counts for the coloring with f(A)=0 for every A.

Again, I then shamelessly changed my code to figure out G(n)=2^n, and so

T(n)=C(n,0)2^n+C(n,1)2^(n-1)+C(n,2)2^(n-2)+...+C(n,n)2^0+1=(1+2)^n+1=3^n+1

Now the rigorous proof will strengthen the arguments and show G(n)=2^n. As always, it appears much less intuitive when we make it formal.

Proof:

Throughout the proof we ignore the trivial coloring with all subsets labeled red, and will add it back at the last step. We're also only concerned with coloring that makes (*) hold. We proceed with a series of lemmas, some of which self-evident.

Lemma 1. f(A)>0 iff A contains a blue subset.

Lemma 2. There is an unique minimum blue subset S.

To prove it, consider two different minimum blue subsets S and S' with same size. Then f(S intersect S') > 0. Since S intersect S' is smaller than S and S', by Lemma 1 this is a contradiction.

Lemma 3. f(A)>0 iff A contains the minimum blue subset S.

If f(A)>0 then B=S intersect A also has f(B)>0, meaning B has a blue subset. Now if A doesn't contain S, then |B|<|S|, a contradiction.

Lemma 4. Let T(n) be the answer to the original problem, and G(n) be the number
of colorings with f({})=1. T(n)=C(n,0)G(n)+C(n,1)G(n-1)+...+C(n,n)G(0)+1

Let the minimum blue subset be S. By Lemma 3 all subsets not containing S should be red. We can check that (1) the problem of coloring subsets containing S is equivalent to the original problem on set N - S when we map each X to X - S. (2) (*) still holds whether A and B contains S or not. So by iterating through all possible S we get the T(n) as described.

Finally, it suffices to show that G(n)=2^n. This is the most interesting part to me.

The idea is to prove that each coloring of singletons {1}, {2}, {3}, ... , {n} leads to one and only one possible colorings of the rest, and therefore G(n)=2^n. 

To show there's only one possible coloring given a coloring of all singletons, first observe that f({i}) is either 1 or 2. So f(doubleton) is either 1, 2, or 4. By induction on |A| we can see that f(A) can take at most one value, which is a power of 2, because f(A)=f({x})f(A-{x})/f({})=f({x})f(A-{x}) where x belongs to A. Now the last remaining piece is to show such f() makes (*) hold all the time. 

Because the only possible f(A) is constructed as product of other f() which are power of 2, L(A)=log2(f(A)) is always integer, and L(A)=L({x})+L(A-{x}). Therefore L(A) is the number of blue singletons contained by A! Now how are non-singletons colored so that f(A) is also the number of A's blue subsets? We can color A blue iff all singletons of A are blue. Then f'(A), the function defined based on this coloring, is the number of A's subsets containing only singletons that are blue, i.e.  f'(A)=2^L(A)=f(A), so this only possible f(A) is also achievable.

Q.E.D.

An example of coloring with f({})=1 when n=4. Notation: color/f/L

{}: B/1/0

{1}: B/2/1    {2}: R/1/0    {3}: R/1/0    {4}: B/2/1

{1,2}: R/2/1    {1,3}: R/2/1    {1,4}: B/4/2    {2,3}: R/1/0    {2,4}: R/2/1    {3,4}: R/2/1

{1,2,3}: R/2/1    {1,2,4}: R/4/2    {1,3,4}: R/4/2    {2,3,4}: R/2/1

{1,2,3,4}: R/4/2
