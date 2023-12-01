+++
title = "Listen"
weight = 1
+++
# Duden/Listen functions
<details>
<summary><h2>Leere_Z</h2></summary>
<ul>
<pre>
Löscht alle Zahlen aus der gegebenen Zahlen Liste.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Zahlen Listen Referenz</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Fügt eine Zahl am Ende der gegeben Zahlen Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_append_int liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenZ</h2></summary>
<ul>
<pre>
Fügt eine Zahl vor einem Index in der gegebenen Zahlen Liste ein.
Es wird nicht überprüft ob der Index valide ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_int liste (index minus 1) elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_Z</h2></summary>
<ul>
<pre>
Fügt eine Zahlen Liste vor einem Index in der gegebenen Zahlen Liste ein.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahlen Liste</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_range_int liste (index minus 1) range (die Größe von 0).

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_Z</h2></summary>
<ul>
<pre>
Fügt eine Zahl am Anfang der gegeben Zahlen Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
[Speichere elm verkettet mit liste in liste.]
efficient_list_prepend_int liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Z</h2></summary>
<ul>
<pre>
Entfernt die Zahl an dem gegeben Index aus der gegeben Zahlen Liste.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere index minus 1 in index.
efficient_list_delete_range_int liste index index (die Größe von 0).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_Z</h2></summary>
<ul>
<pre>
Entfernt alle Zahlen aus der Liste im Bereich [start, end] (inklusiv)
</pre>
	<li>Parameters: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_delete_range_int liste (start minus 1) (end minus 1) (die Größe von 0).

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_Z</h2></summary>
<ul>
<pre>
Füllt die gegebene Zahlen Liste mit der gegebenen Zahl.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Zahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt den Index der gegebenen Zahl aus der Liste zurück oder -1 wenn die Zahl nicht in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Zahlen Liste</code>, <code>Zahl</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Zahl in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Zahlen Liste</code>, <code>Zahl</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Zahlen Liste leer ist.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Zahlen Liste</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
Löscht alle Kommazahlen aus der gegebenen Kommazahlen Liste.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Kommazahlen Listen Referenz</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Fügt eine Kommazahl am Ende der gegeben Kommazahlen Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Kommazahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_append_float liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenK</h2></summary>
<ul>
<pre>
Fügt eine Kommazahl vor einem Index in der gegebenen Kommazahlen Liste ein.
Es wird nicht überprüft ob der Index valide ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code>, <code>Kommazahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_float liste (index minus 1) elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_K</h2></summary>
<ul>
<pre>
Fügt eine Kommazahlen Liste vor einem Index in der gegebenen Kommazahlen Liste ein.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code>, <code>Kommazahlen Liste</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_range_float liste (index minus 1) range (die Größe von 0,0).

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_K</h2></summary>
<ul>
<pre>
Fügt eine Kommazahl am Anfang der gegeben Kommazahlen Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Kommazahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
[Speichere elm verkettet mit liste in liste.]
efficient_list_prepend_float liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_K</h2></summary>
<ul>
<pre>
Entfernt die Kommazahl an dem gegeben Index aus der gegeben Kommazahlen Liste.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere index minus 1 in index.
efficient_list_delete_range_float liste index index (die Größe von 0,0).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_K</h2></summary>
<ul>
<pre>
Entfernt alle Kommazahlen aus der Liste im Bereich [start, end] (inklusiv)
</pre>
	<li>Parameters: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_delete_range_float liste (start minus 1) (end minus 1) (die Größe von 0,0).

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_K</h2></summary>
<ul>
<pre>
Füllt die gegebene Kommazahlen Liste mit der gegebenen Kommazahl.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Kommazahlen Listen Referenz</code>, <code>Kommazahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt den Index der gegebenen Kommazahl aus der Liste zurück oder -1 wenn die Kommazahl nicht in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Kommazahl in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Kommazahlen Liste leer ist.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Kommazahlen Liste</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
Löscht alle Booleans aus der gegebenen Boolean Liste.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Boolean Listen Referenz</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Fügt einen Boolean am Ende der gegeben Boolean Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Boolean</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_append_bool liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenB</h2></summary>
<ul>
<pre>
Fügt einen Boolean vor einem Index in der gegebenen Boolean Liste ein.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Zahl</code>, <code>Boolean</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_bool liste (index minus 1) elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_B</h2></summary>
<ul>
<pre>
Fügt eine Boolean Liste vor einem Index in der gegebenen Boolean Liste ein.
Es wird nicht überprüft ob der Index valide ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Zahl</code>, <code>Boolean Liste</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_range_bool liste (index minus 1) range (die Größe von wahr).

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_B</h2></summary>
<ul>
<pre>
Fügt einen Boolean am Anfang der gegeben Boolean Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Boolean</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
[Speichere elm verkettet mit liste in liste.]
efficient_list_prepend_bool liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_B</h2></summary>
<ul>
<pre>
Entfernt den Boolean an dem gegeben Index aus der gegeben Boolean Liste.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere index minus 1 in index.
efficient_list_delete_range_bool liste index index (die Größe von wahr).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_B</h2></summary>
<ul>
<pre>
Entfernt alle Booleans aus der Liste im Bereich [start, end] (inklusiv)
</pre>
	<li>Parameters: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_delete_range_bool liste (start minus 1) (end minus 1) (die Größe von wahr).

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_B</h2></summary>
<ul>
<pre>
Füllt die gegebene Kommazahlen Liste mit der gegebenen Kommazahl.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Boolean Listen Referenz</code>, <code>Boolean</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt den Index des gegebenen Booleans aus der Liste zurück oder -1 wenn der Boolean nicht in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Boolean Liste</code>, <code>Boolean</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob der Boolean in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Boolean Liste</code>, <code>Boolean</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Boolean Liste leer ist.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Boolean Liste</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
Löscht alle Buchstaben aus der gegebenen Buchstaben Liste.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Buchstaben Listen Referenz</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Fügt einen Buchstaben am Ende der gegeben Buchstaben Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_append_char liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenC</h2></summary>
<ul>
<pre>
Fügt einen Buchstaben vor einem Index in der gegebenen Buchstaben Liste ein.
Es wird nicht überprüft ob der Index valide ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_char liste (index minus 1) elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_C</h2></summary>
<ul>
<pre>
Fügt eine Buchstaben Liste vor einem Index in der gegebenen Buchstaben Liste ein.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code>, <code>Buchstaben Liste</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_range_char liste (index minus 1) range (die Größe von ' ').

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_C</h2></summary>
<ul>
<pre>
Fügt einen Buchstaben am Anfang der gegeben Buchstaben Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
[Speichere elm verkettet mit liste in liste.]
efficient_list_prepend_char liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_C</h2></summary>
<ul>
<pre>
Entfernt den Buchstaben an dem gegeben Index aus der gegeben Buchstaben Liste.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere index minus 1 in index.
efficient_list_delete_range_char liste index index (die Größe von ' ').

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_C</h2></summary>
<ul>
<pre>
Entfernt alle Buchstaben aus der Liste im Bereich [start, end] (inklusiv)
</pre>
	<li>Parameters: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_delete_range_char liste (start minus 1) (end minus 1) (die Größe von ' ').

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_C</h2></summary>
<ul>
<pre>
Füllt die gegebene Buchstaben Liste mit dem gegebenen Buchstaben.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Buchstaben Listen Referenz</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt den Index des gegebenen Buchstabens aus der Liste zurück oder -1 wenn der Buchstabe nicht in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Buchstaben Liste</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob der Buchstabe in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Buchstaben Liste</code>, <code>Buchstabe</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Buchstaben Liste leer ist.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Buchstaben Liste</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
Löscht alle Texte aus der gegebenen Text Liste.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Text Listen Referenz</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Fügt einen Text am Ende der gegeben Text Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Text</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
	<li><code>&#34;füge &lt;elm&gt; an &lt;liste&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_append_string liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>EinfügenT</h2></summary>
<ul>
<pre>
Fügt einen Text vor einem Index in der gegebenen Text Liste ein.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Zahl</code>, <code>Text</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_string liste (index minus 1) elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Einfügen_Bereich_T</h2></summary>
<ul>
<pre>
Fügt eine Text Liste vor einem Index in der gegebenen Text Liste ein.
Es wird nicht überprüft ob der Index valide ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code>, <code>range</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Zahl</code>, <code>Text Liste</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Setze die Elemente in &lt;range&gt; an die Stelle &lt;index&gt; von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_insert_range_string liste (index minus 1) range (die Größe von "").

</code>
</pre>
</details>

<details>
<summary><h2>Voranstellen_T</h2></summary>
<ul>
<pre>
Fügt einen Text am Anfang der gegeben Text Liste hinzu.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Text</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
[Speichere elm verkettet mit liste in liste.]
efficient_list_prepend_string liste elm (die Größe von elm).

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_T</h2></summary>
<ul>
<pre>
Entfernt den Text an dem gegeben Index aus der gegeben Text Liste.
</pre>
	<li>Parameters: <code>liste</code>, <code>index</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere index minus 1 in index.
efficient_list_delete_range_string liste index index (die Größe von "").

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Bereich_T</h2></summary>
<ul>
<pre>
Entfernt alle Texte aus der Liste im Bereich [start, end] (inklusiv)
</pre>
	<li>Parameters: <code>liste</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
efficient_list_delete_range_string liste (start minus 1) (end minus 1) (die Größe von "").

</code>
</pre>
</details>

<details>
<summary><h2>Füllen_T</h2></summary>
<ul>
<pre>
Füllt die gegebene Text Liste mit dem gegebenen Text.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Text Listen Referenz</code>, <code>Text</code></li>
	<li>Return type: <code>nichts</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt den Index des gegebenen Textes aus der Liste zurück oder -1 wenn der Text nicht in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Text Liste</code>, <code>Text</code></li>
	<li>Return type: <code>Zahl</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob der Text in der Liste vorhanden ist.
</pre>
	<li>Parameters: <code>liste</code>, <code>elm</code></li>
	<li>Parameter types: <code>Text Liste</code>, <code>Text</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Gibt zurück ob die Text Liste leer ist.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Text Liste</code></li>
	<li>Return type: <code>Boolean</code></li>
</ul>

<h3>Aliases</h3>
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
<pre>
Verkettet alle Buchstaben der gegebenen Liste zu einem Text und gibt diesen zurück.
</pre>
	<li>Parameters: <code>liste</code></li>
	<li>Parameter type: <code>Buchstaben Liste</code></li>
	<li>Return type: <code>Text</code></li>
</ul>

<h3>Aliases</h3>
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


