# Solution
## Problem
Given an array of numbers with duplicated nums :$A = \left \{ n_{0}... \right \}$
Given a target number: $t$
Given a target length: $y$

Find all the set $S_{y}^{t}$ which stasfied the following condition:
1. $\left | S_{y}^{t} \right | = y$ 
2. $  s \subseteqq S_{y}^{t} \rightarrow s \subseteqq A $ 
3. $ \sum S_{y}^{t} = t $

## Solution
1. ##### Recursive

##### Train of thought:

Given a number $n_{x}$ :
$$ 
S_{y}^{t}=\left\{ n_{x} \right\} \bigcup S_{y-1}^{t-n_{x}}(eq1)
$$

So we can get a recursive function $eq2$:
1. $F(y,t)=x*F(y-1,t-x)\;\;(y>1)\;\;(for\;unique(x)\;in\;A)$
2. $F(y,t)=t\;\; (y=1)$

However this function will get duplicated sets
This can be prevented by sort the set while selecting elements from $A$
Which means:
if we selected $x_{i_{0}}$ ($i_{0}$ represents the order of the element in the set) $\vee x_{i}<x_{i_{0}}(i>i_{0})$

##### Proof:
1. Prove the set from the recursive function meet the requirements and are distinct:
a. $\vee S_{1},S_{2}\;s_{i}^{1}\neq s_{i}^{2}$($i$ represent the order in the set. This is guaraenteed by iterating through unique element on each selection )
b. Given $ S $ is orderd. if $  S_{1}=S_{2}$ then $ s_{i}^{1}=s_{i}^{2}$ which conflicts with a. Thus $\vee S_{i} \neq S_{j}$
c. $\because eq1 \therefore \vee \sum S = t $
2. Prove all the $S$ are generated from $eq2$:
a. if $\exists S_{0} $ not generated from $eq2$
&nbsp;&nbsp;&nbsp;&nbsp;then $\exists s_{i}^{0} \nsubseteq A$
&nbsp;&nbsp;&nbsp;&nbsp;$\because$ all the numbers is selected from $A$. $\therefore \nexists S_{0}$