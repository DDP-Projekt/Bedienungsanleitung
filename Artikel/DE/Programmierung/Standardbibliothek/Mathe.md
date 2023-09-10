# Duden/Mathe Funktionen
<details>
<summary><h2>PI</h2></summary>
<ul>
<pre>
Gibt den Wert von der Kreiszahl "PI" (π) mit 15 Nachkommastellen zurück:
3,141592653589793
</pre>
</li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;PI&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib 3,141592653589793 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>E</h2></summary>
<ul>
<pre>
Gibt den Wert der Eulerschen Zahl "E" mit 15 Nachkommastellen zurück:
2,718281828459045
</pre>
</li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;E&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib 2,718281828459045 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>TAU</h2></summary>
<ul>
<pre>
Gibt den Wert der alternativen Kreiszahl "TAU" (τ) mit 15 Nachkommastellen zurück:
6,283185307179586
Es entspricht exakt den Wert von 2 mal PI.
</pre>
</li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;TAU&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib 6,283185307179586 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>PHI</h2></summary>
<ul>
<pre>
Gibt den Wert des Goldenen Schittes "PHI" (Φ) mit 15 Nachkommastellen zurück:
1,618033988749895
</pre>
</li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;PHI&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib 1,618033988749895 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>max</h2></summary>
<ul>
<pre>
Wenn a >= b ist wird a zurück gegeben.
Wenn a < b ist wird b zurückgegeben.
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die größere Zahl von &lt;a&gt; und &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn a größer als, oder b ist, gib a zurück.
Gib b zurück.

</code>
</pre>
</details>

<details>
<summary><h2>min</h2></summary>
<ul>
<pre>
Wenn a <= b ist wird a zurückgegeben.
Wenn a > b ist wird b zurückgegeben.
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die kleinere Zahl von &lt;a&gt; und &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn a kleiner als, oder b ist, gib a zurück.
Gib b zurück.

</code>
</pre>
</details>

<details>
<summary><h2>clamp</h2></summary>
<ul>
<pre>
Wenn wert > max ist, wird max zurückgegeben.
Wenn wert < min ist, wird min zurückgegeben.
Wenn min < wert < max ist, wird wert zurückgegeben.
</pre>
	<li>Parameter: <code>wert</code>, <code>max</code>, <code>min</code></li>
	<li>Parameter Typen: <code>Zahl</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;wert&gt; zwischen &lt;min&gt; und &lt;max&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn wert größer als max ist, gib max zurück.
Wenn wert kleiner als min ist, gib min zurück.
Gib wert zurück.

</code>
</pre>
</details>

<details>
<summary><h2>maxK</h2></summary>
<ul>
<pre>
Wenn a >= b ist wird a zurück gegeben.
Wenn a < b ist wird b zurückgegeben.
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Kommazahl</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die größere Zahl von &lt;a&gt; und &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn a größer als, oder b ist, gib a zurück.
Gib b zurück.

</code>
</pre>
</details>

<details>
<summary><h2>minK</h2></summary>
<ul>
<pre>
Wenn a <= b ist wird a zurückgegeben.
Wenn a > b ist wird b zurückgegeben.
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Kommazahl</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die kleinere Zahl von &lt;a&gt; und &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn a kleiner als, oder b ist, gib a zurück.
Gib b zurück.

</code>
</pre>
</details>

<details>
<summary><h2>clampK</h2></summary>
<ul>
<pre>
Wenn wert > max ist, wird max zurückgegeben.
Wenn wert < min ist, wird min zurückgegeben.
Wenn min < wert < max ist, wird wert zurückgegeben.
</pre>
	<li>Parameter: <code>wert</code>, <code>max</code>, <code>min</code></li>
	<li>Parameter Typen: <code>Kommazahl</code>, <code>Kommazahl</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;wert&gt; zwischen &lt;min&gt; und &lt;max&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn wert größer als max ist, gib max zurück.
Wenn wert kleiner als min ist, gib min zurück.
Gib wert zurück.

</code>
</pre>
</details>

<details>
<summary><h2>sign</h2></summary>
<ul>
<pre>
Wenn wert < 0 ist, wird -1 zurückgegeben.
Wenn wert > 0 ist, wird 1 zurückgegeben.
Wenn wert = 0 ist, wird 0 zurückgegeben.
</pre>
	<li>Parameter: <code>wert</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;das Vorzeichen von &lt;wert&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn wert kleiner als 0 ist, gib -1 zurück.
Wenn aber wert größer als 0 ist, gib 1 zurück.
Gib 0 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>signK</h2></summary>
<ul>
<pre>
Wenn wert < 0 ist, wird -1 zurückgegeben.
Wenn wert > 0 ist, wird 1 zurückgegeben.
Wenn wert = 0 ist, wird 0 zurückgegeben.
</pre>
	<li>Parameter: <code>wert</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;das Vorzeichen von &lt;wert&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn wert kleiner als 0 ist, gib -1 zurück.
Wenn aber wert größer als 0 ist, gib 1 zurück.
Gib 0 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>floor</h2></summary>
<ul>
<pre>
Rundet wert nach unten.
</pre>
	<li>Parameter: <code>wert</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;wert&gt; nach unten gerundet&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib wert minus (wert minus wert als Zahl) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>ceil</h2></summary>
<ul>
<pre>
Rundet wert nach oben.
</pre>
	<li>Parameter: <code>wert</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;wert&gt; nach oben gerundet&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib wert plus (1 minus (wert minus wert als Zahl)) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>trunc</h2></summary>
<ul>
<pre>
Schneidet alle Kommastellen von wert ab.
</pre>
	<li>Parameter: <code>wert</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;wert&gt; trunkiert&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib (wert als Zahl) als Kommazahl zurück.

</code>
</pre>
</details>

<details>
<summary><h2>sin</h2></summary>
<ul>
<pre>
Berechnet den Sinus von v.
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Sinus von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Sinus von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>cos</h2></summary>
<ul>
<pre>
Berechnet den Kosinus von v.
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Kosinus von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Kosinus von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>tan</h2></summary>
<ul>
<pre>
Berechnet den Tangens von v.
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Tangens von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Tangens von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>asin</h2></summary>
<ul>
<pre>
Berechnet den Arkussinus von v. (sin⁻¹)
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Arkussinus von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Arkussinus von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>acos</h2></summary>
<ul>
<pre>
Berechnet den Arkuskosinus von v. (cos⁻¹)
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Arkuskosinus von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Arkuskosinus von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>atan</h2></summary>
<ul>
<pre>
Berechnet den Arkustangens von v. (tan⁻¹)
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Arkustangens von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Arkustangens von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>sinh</h2></summary>
<ul>
<pre>
Berechnet den Hyperbelsinus von v.
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Hyperbelsinus von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Hyperbelsinus von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>cosh</h2></summary>
<ul>
<pre>
Berechnet den Hyperbelkosinus von v.
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Hyperbelkosinus von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Hyperbelkosinus von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>tanh</h2></summary>
<ul>
<pre>
Berechnet den Hyperbeltangens von v.
</pre>
	<li>Parameter: <code>v</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Hyperbeltangens von &lt;v&gt;&#34;</code></li>
	<li><code>&#34;den Hyperbeltangens von &lt;v&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
Implementiert in <code>"libddpstdlib.a"</code>
</details>

<details>
<summary><h2>groesster_gemeinsamer_Teiler</h2></summary>
<ul>
<pre>
Eine Funktion, die den größten gemeinsamen Teiler zweier Zahlen, <a> und <b>, als Zahl zurück gibt.
Zeitkomplexität: O(n)
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der größte gemeinsame Teiler von &lt;a&gt; und &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl t ist 0.
Solange b ungleich 0 ist, mache:
	Speichere b in t.
	Speichere (a modulo t) in b.
	Speichere t in a.
Gib a zurück.

</code>
</pre>
</details>

<details>
<summary><h2>kleinster_gemeinsamer_Teiler</h2></summary>
<ul>
<pre>
Eine Funktion, die den kleinsten gemeinsamen Teiler zweier Zahlen, <a> und <b>, als Zahl zurück gibt.
Zeitkomplexität: O(n)
</pre>
	<li>Parameter: <code>a</code>, <code>b</code></li>
	<li>Parameter Typen: <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;das kleinste gemeinsame Vielfache von &lt;a&gt; und &lt;b&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib (der Betrag von (a mal b) durch (der größte gemeinsame Teiler von a und b)) als Zahl zurück.

</code>
</pre>
</details>

<details>
<summary><h2>ist_teilbar</h2></summary>
<ul>
<pre>
Ob der divident durch den divisor teilbar ist.
Auch: Ob divident modulo divisor = 0 ist.
</pre>
	<li>Parameter: <code>dividend</code>, <code>divisor</code></li>
	<li>Parameter Typen: <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;dividend&gt; durch &lt;divisor&gt; teilbar ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib dividend modulo divisor gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Primfaktorzerlegung</h2></summary>
<ul>
<pre>
Eine Funktion, die eine Zahlen Liste von allen Primfaktoren der Zahl <z> gibt.  
Zeitkomplexität: O(sqrt(n))
</pre>
	<li>Parameter: <code>z</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahlen Liste</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Primfaktoren von &lt;z&gt;&#34;</code></li>
	<li><code>&#34;alle Primfaktoren von &lt;z&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahlen Liste faktoren ist eine leere Zahlen Liste.

Solange z durch 2 teilbar ist, mache:
	Speichere faktoren verkettet mit 2 in faktoren.
	Speichere (z durch 2) als Zahl in z.

Die Zahl i ist 3.
Solange i kleiner als, oder die 2. Wurzel von z als Zahl ist, mache:
	Solange z durch i teilbar ist, mache:
		Speichere faktoren verkettet mit i in faktoren.
		Speichere (z durch i) als Zahl in z.
	Erhöhe i um 2.

Wenn z größer als 2 ist, dann:
	Speichere faktoren verkettet mit z in faktoren.

Gib faktoren zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Teilerzerlegung</h2></summary>
<ul>
<pre>
Gibt eine Zahlen Liste von alle Zahlen, die durch <z> geteilt werden können.
Zeitkomplexität: O(n)
</pre>
	<li>Parameter: <code>z</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahlen Liste</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;alle Teiler von &lt;z&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahlen Liste teiler ist eine leere Zahlen Liste.

Für jede Zahl i von z bis 1 mit Schrittgröße -1, mache:
	Wenn z durch i teilbar ist, speichere teiler verkettet mit i in teiler.	

Gib teiler zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Quadriere</h2></summary>
<ul>
<pre>
Quadriert (hoch 2) die gegebene Zahl.
</pre>
	<li>Parameter: <code>x</code></li>
	<li>Parameter Typ: <code>Kommazahlen Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Quadriere &lt;x&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Vervielfache x um x.

</code>
</pre>
</details>

<details>
<summary><h2>Quadriere_Wert</h2></summary>
<ul>
<pre>
Gibt die gegebene Zahl hoch 2 zurück.
</pre>
	<li>Parameter: <code>x</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;x&gt; zum quadrat&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib x mal x zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Ganze_Zahl</h2></summary>
<ul>
<pre>
Gibt zurück ob die gegebene Kommazahl eine ganze Zahl ist. (...; -2,0; -1,0; 0,0; 1,0; 2,0; ...)
</pre>
	<li>Parameter: <code>x</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;x&gt; eine ganze Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib (x als Zahl) als Kommazahl gleich x ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Gerade_Zahl</h2></summary>
<ul>
<pre>
Gibt zurück ob die gegebene Zahl eine gerade Zahl ist. (x mod 2 = 0)
</pre>
	<li>Parameter: <code>x</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;x&gt; eine gerade Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib x modulo 2 gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Gerade_Kommazahl</h2></summary>
<ul>
<pre>
Gibt zurück ob die gegebene Kommazahl eine gerade Zahl ist. ((int)x mod 2 = 0)
</pre>
	<li>Parameter: <code>x</code></li>
	<li>Parameter Typ: <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;x&gt; eine gerade Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib x als Zahl modulo 2 gleich 0 ist zurück.

</code>
</pre>
</details>


