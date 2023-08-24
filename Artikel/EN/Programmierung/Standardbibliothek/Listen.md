# Duden/Listen Funktionen
<details>
<summary><h2>Leere_Z</h2></summary>
<ul>
<pre>
Zahlen Listen Funktionen
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Zahlen Listen Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Leere &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere eine leere Zahlen Liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Hinzufüge_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste verkettet mit elm in liste.

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenZ</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahlen Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste gleich 0 ist, verlasse die Funktion.

Wenn index gleich 1 ist und die Länge von liste größer als 1 ist, dann:
	Speichere eine leere Zahlen Liste in liste.
Wenn aber index gleich 1 ist, dann:
	Speichere liste von 2 bis (die Länge von liste) in liste.
Wenn aber index gleich die Länge von liste ist, dann:
	Speichere liste von 1 bis (die Länge von liste minus 1) in liste.
Sonst:
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Fülle &lt;liste&gt; mit &lt;elm&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahlen Liste neueListe ist die Länge von liste Mal elm.
Speichere neueListe in liste.

</code>
</pre>
</details>

<details>
<summary><h2>IndexVon_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Zahlen Liste</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Index von &lt;elm&gt; in &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Zahl i von 1 bis (die Länge von liste), wenn liste an der Stelle i gleich elm ist, gib i zurück.
Gib -1 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Enthält_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Zahlen Liste</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; &lt;elm&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Zahl z in liste, wenn z gleich elm ist, gib wahr zurück.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstLeer_Z</h2></summary>
<ul>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Zahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; leer ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Leere_K</h2></summary>
<ul>
<pre>
Kommazahlen Listen Funktionen
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Listen Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Leere &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere eine leere Kommazahlen Liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Hinzufüge_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste verkettet mit elm in liste.

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenK</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste gleich 0 ist, verlasse die Funktion.

Wenn index gleich 1 ist und die Länge von liste größer als 1 ist, dann:
	Speichere eine leere Kommazahlen Liste in liste.
Wenn aber index gleich 1 ist, dann:
	Speichere liste von 2 bis (die Länge von liste) in liste.
Wenn aber index gleich die Länge von liste ist, dann:
	Speichere liste von 1 bis (die Länge von liste minus 1) in liste.
Sonst:
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Kommazahlen Listen Referenz</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Fülle &lt;liste&gt; mit &lt;elm&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahlen Liste neueListe ist die Länge von liste Mal elm.
Speichere neueListe in liste.

</code>
</pre>
</details>

<details>
<summary><h2>IndexVon_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Index von &lt;elm&gt; in &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
Gib -1 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Enthält_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; &lt;elm&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Kommazahl z in liste, wenn z gleich elm ist, gib wahr zurück.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstLeer_K</h2></summary>
<ul>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; leer ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Leere_B</h2></summary>
<ul>
<pre>
Boolean Listen Funktionen
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Boolean Listen Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Leere &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere eine leere Boolean Liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Hinzufüge_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste verkettet mit elm in liste.

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenB</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Zahl</code>, <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Zahl</code>, <code>Boolean Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste gleich 0 ist, verlasse die Funktion.

Wenn index gleich 1 ist und die Länge von liste größer als 1 ist, dann:
	Speichere eine leere Boolean Liste in liste.
Wenn aber index gleich 1 ist, dann:
	Speichere liste von 2 bis (die Länge von liste) in liste.
Wenn aber index gleich die Länge von liste ist, dann:
	Speichere liste von 1 bis (die Länge von liste minus 1) in liste.
Sonst:
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Boolean Listen Referenz</code>, <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Fülle &lt;liste&gt; mit &lt;elm&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Boolean Liste neueListe ist die Länge von liste Mal elm.
Speichere neueListe in liste.

</code>
</pre>
</details>

<details>
<summary><h2>IndexVon_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Boolean Liste</code>, <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Index von &lt;elm&gt; in &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
Gib -1 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Enthält_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Boolean Liste</code>, <code>Boolean</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; &lt;elm&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jeden Boolean b in liste, wenn b gleich elm ist, gib wahr zurück.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstLeer_B</h2></summary>
<ul>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Boolean Liste</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; leer ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Leere_C</h2></summary>
<ul>
<pre>
Buchstaben Listen Funktionen
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Buchstaben Listen Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Leere &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere eine leere Buchstaben Liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Hinzufüge_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste verkettet mit elm in liste.

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenC</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code>, <code>Buchstaben Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste gleich 0 ist, verlasse die Funktion.

Wenn index gleich 1 ist und die Länge von liste größer als 1 ist, dann:
	Speichere eine leere Buchstaben Liste in liste.
Wenn aber index gleich 1 ist, dann:
	Speichere liste von 2 bis (die Länge von liste) in liste.
Wenn aber index gleich die Länge von liste ist, dann:
	Speichere liste von 1 bis (die Länge von liste minus 1) in liste.
Sonst:
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Buchstaben Listen Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Fülle &lt;liste&gt; mit &lt;elm&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Buchstaben Liste neueListe ist die Länge von liste Mal elm.
Speichere neueListe in liste.

</code>
</pre>
</details>

<details>
<summary><h2>IndexVon_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Buchstaben Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Index von &lt;elm&gt; in &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
Gib -1 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Enthält_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Buchstaben Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; &lt;elm&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jeden Buchstaben b in liste, wenn b gleich elm ist, gib wahr zurück.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstLeer_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Buchstaben Liste</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; leer ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Leere_T</h2></summary>
<ul>
<pre>
Text Listen Funktionen
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Text Listen Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Leere &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere eine leere Text Liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Hinzufüge_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste verkettet mit elm in liste.

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenT</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Zahl</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit elm verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Zahl</code>, <code>Text Liste</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (index minus 1) verkettet mit range verkettet mit liste von index  bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit liste in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>index</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste gleich 0 ist, verlasse die Funktion.

Wenn index gleich 1 ist und die Länge von liste größer als 1 ist, dann:
	Speichere eine leere Text Liste in liste.
Wenn aber index gleich 1 ist, dann:
	Speichere liste von 2 bis (die Länge von liste) in liste.
Wenn aber index gleich die Länge von liste ist, dann:
	Speichere liste von 1 bis (die Länge von liste minus 1) in liste.
Sonst:
	Speichere liste von 1 bis (index minus 1) verkettet mit liste von (index plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere liste von 1 bis (start minus 1) verkettet mit liste von (end plus 1) bis (die Länge von liste) in liste.

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Listen Referenz</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Fülle &lt;liste&gt; mit &lt;elm&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Text Liste neueListe ist die Länge von liste Mal elm.
Speichere neueListe in liste.

</code>
</pre>
</details>

<details>
<summary><h2>IndexVon_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Index von &lt;elm&gt; in &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jede Zahl i von 1 bis (die Länge von liste), Wenn liste an der Stelle i gleich elm ist, Gib i zurück.
Gib -1 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Enthält_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Liste</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; &lt;elm&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jeden Text t in liste, wenn t gleich elm ist, gib wahr zurück.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstLeer_T</h2></summary>
<ul>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Text Liste</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; leer ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib [wahr wenn] die Länge von liste gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Aneinandergehängt_C</h2></summary>
<ul>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Buchstaben Liste</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; aneinandergehängt&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text t ist "".
Für jeden Buchstabe b in liste, mache:
	Speichere t verkettet mit b in t.
Gib t zurück.

</code>
</pre>
</details>


