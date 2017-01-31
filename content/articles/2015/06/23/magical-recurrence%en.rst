[Algorithm] Magical Recurrence
##############################

:date: 2015-06-23
:author: Shen-Fu Tsai
:tags: Algorithm,Math
:category: Algorithm

In a recent algorithmic coding contest which I didn't do well, the `hardest problem`_ killed me. It distinguished me from other superior coders. But it's still an interesting one. The problem is essentially to solve a 0/1 knapsack problem on a tree where each node is associated with an item with a given value and size. The only constraint is that the set of selected nodes should be connected.

Noting that the tree is not rooted, it differs from the knapsack problem on a rooted tree, i.e. TPK, in only that the root has to be chosen in TPK, if possible. So one of the insights I didn't come up with in order to accommodate TPK is to iteratively find and use the "central" node of the tree as root, and solve the rooted sub-problem recursively. Then, if it takes :math:`O(N)` to solve a TPK for a tree with :math:`N` nodes, the overall complexity will be :math:`O(N\log(N))` instead of :math:`N^2`. Gladly I came up with my own :math:`O(N)` algorithm to find central node after the contest, so it is a good exercise to think of one by yourself. To clarify, a central node has none of its child subtrees sized more than :math:`N/2`.

The main reason I write this article is the algorithm that solves TPK in :math:`O(Nh)`, where :math:`h` is the bag capacity. During contest I did find a paper describing a DFS-based DP algorithm, which I adapted to keep :math:`D[u][k][b]`, meaning the maximum value achievable with capacity :math:`b` if every node on the path from root to node u are chosen, and all other nodes up to :math:`u` in DFS traversal order, plus the the :math:`k`-th child subtree of :math:`u` are considered.  When :math:`k` is 0 no child subtree of :math:`u` is considered. For anyone interested, it is described in [1]. I'm guessing the better solution we're going to discuss is covered by [2], but I didn't manage to get its content for free in the course of 2 days. The least I'd have to pay was $24. Regardless whether [2] is the key to the solution or not I've at least figured something out after the contest.

What is interesting is the `solution`_ posted after the contest and likely covered in [2] is to keep :math:`D[u][b]` to indicate the maximum value attainable up to node :math:`u` in the DFS traversal order with capacity :math:`b`.  The recurrence is surprisingly simple:

Recurence
+++++++++

(1) if (:math:`size[u]\leq b`) { :math:`f[u][b]=MAX(f[u-1][b-size[u]]+value[u], f[u][b])`; }
(2) :math:`f[u+subtree\_size[u]-1][b]=MAX(f[u+subtree\_size[u]-1][b], f[u-1][b])`;

Here :math:`u-1` is the node traversed right before :math:`u`.

Does the recurrence make sense? The second update is very obvious -- if we don't take node :math:`u`, then the whole subtree rooted at :math:`u` is not taken. However, the first update doesn't sound correct to me -- if we take node :math:`u`, then we have to make sure all ancestors of :math:`u` are taken too, but :math:`f[u-1][b-size[i]]` doesn't imply it, which is simply the optimal value obtainable when considering up to the node before :math:`u`!

With confusion, I created a simple test case and check :math:`D` values for some intermediate nodes. In particular I created a first tree and appended one more node at the end of its DFS traversal to get the second tree. Then I ran the algorithm, and found the :math:`D` values for a node present in both tree DIFFER!

So obviously the interpretation of :math:`D` by the problem setter is NOT precise, if his/her solution actually works. Eventually, I'm able to show this magical recurrence does work.

Lemma
+++++
If :math:`u` is the last traversed node by DFS, then :math:`D[u][b]` is the optimal value that could be collected from the tree with bag capacity :math:`b`.

Proof
+++++

We will show by induction on :math:`N`, the tree size, as induction are close to DP in nature.  Suppose the Lemma holds for every tree with up to :math:`N-1` nodes. First, the case that the root has only one child is trivial -- we could descend from root until we hit a node with multiple children and reduce the problem. So assume the root has multiple children, and consider the last child :math:`u` of root traversed by DFS. Let :math:`t(u)` be the subtree rooted at :math:`u`. Let :math:`T` be the original tree and by removing :math:`t(u)` from :math:`T` we get :math:`T_1` which ends at :math:`u-1`, the node visited by DFS right before :math:`u`.

By induction, :math:`D[u-1][b]` is the optimal value for :math:`T_1`. Now the only subtree\_size affected by adding :math:`t(u)` is :math:`subtree\_size[root]`, which doesn't matter to us. So we know :math:`D[x][]` doesn't change for any :math:`x` between the root and :math:`u-1`, inclusive, and we will show :math:`D[v][b]` is the optimal value for :math:`T`, where :math:`v` is the last node in :math:`T`.

How will :math:`D[v][b]` get updated? By (2) in the recurrence it is at least :math:`D[u-1][b]` -- sure, if we don't take :math:`u`, then :math:`D[u-1][b]` is our answer and by induction we know it is correct. In fact this is the only direct (2) update on :math:`t(u)` from :math:`T_1` because by DFS nature no subtree rooted in :math:`T_1` ends in a node in :math:`t(u)`.  The only remaining influence of :math:`T_1` on :math:`t(u)` is the update of :math:`D[u][b]` based on (1).

If :math:`b<size[u]`, (1) couldn't carry out, so all nodes in :math:`t(u)` has no :math:`D` value except :math:`v` which has :math:`D[v][b]=D[u-1][b]`. This is expected, because the whole :math:`t(u)` couldn't be taken since :math:`size[u]` exceeds capacity :math:`b`.

If :math:`b\geq size[u]`, we set :math:`D[u][b]=D[u-1][b-size[u]]+value[u]` directly. Then the DP algorithm computes :math:`D` in the subproblem of :math:`t(u)` as usual except an added term of :math:`D[u-1][b-size[u]]` on :math:`D[u][b]`. All other nodes in :math:`t(u)` are updated as usual except with the blessing of :math:`D[u-1][b-size[u]]`, so we know that before the aformentioned (2) update on :math:`D[v][b]`, this value is the optimal value we can get if the entire tree :math:`T` is considered and node :math:`u` is taken.  After the (2) update, :math:`D[v][b]` also considers not taking :math:`u(t)`. Therefore :math:`D[v][b]` is optimal.

Q.E.D.

By the way, although both DP algorithms have the same asymptotic time complexity, I couldn't make the one in [1] pass the contest's time limit even with quite some optimization! It is because each of its node has to compute one extra value, so the table it keeps is at twice as big. Thus is the power of the magical recurrence!

----

`post <http://oathbystyx.blogspot.com/2015/06/a-magical-recurrence-for-tree-knapsack.html>`_
by
`Shen-Fu Tsai <{filename}/pages/en/sftsai.rst>`_

References

.. [1] `D. S. Johnson and K. A. Niemi. On Knapsacks, partitions, and a new dynamic programming technique for trees <http://www.jstor.org/stable/3689406?seq=1#page_scan_tab_contents>`_

.. [2] `Geon Cho and Dong X. Shaw. A Depth-First Dynamic Programming Algorithm for the Tree Knapsack Problem <http://pubsonline.informs.org/doi/abs/10.1287/ijoc.9.4.431?journalCode=ijoc>`_

.. _hardest problem: https://www.hackerrank.com/contests/w15/challenges/a-knapsack-problem
.. _solution: https://www.hackerrank.com/contests/w15/challenges/a-knapsack-problem/editorial
