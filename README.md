# betriebsrat

**Know your rights at work — for employees and works council members.**

> **[🇩🇪 Deutsche Version weiter unten](#deutsch)**

German labour law is full of rights most people don't know they have. This tool makes them instantly accessible — in plain language, in German or English, without needing to know any law yourself.

---

## What can I ask?

Just describe your situation. Here are examples of what the tool considers when answering:

**As an employee:**

> "Can I get severance? I've worked here 7 years and earn €3,500 a month."

Based on 7 years and €3,500/month, the Munich formula estimates around €12,250. Whether you can claim it depends on whether a Sozialplan was agreed — or bypassed.

> "I received a dismissal notice. Was the works council properly consulted?"

The works council must be heard before every dismissal (§ 102 BetrVG). If the employer skipped this step or gave incomplete information, the dismissal is invalid — regardless of the stated reason.

> "I'm pregnant and my employer wants to end my contract."

Termination during pregnancy is prohibited under the Mutterschutzgesetz — even if the employer doesn't know yet. You have 2 weeks after receiving the notice to inform them, and the protection applies retroactively.

> "My employer introduced monitoring software. Did the works council have to agree?"

Yes — any software that can track performance or behaviour requires works council consent (§ 87 BetrVG). If it was rolled out without agreement, the BR can demand it be stopped.

> "I was transferred to a different location without my agreement. Is that allowed?"

Transfers require works council approval (§ 99 BetrVG). If the BR wasn't consulted, or the new location significantly changes your working conditions, the transfer can be challenged.

---

**As a works council member:**

> "We received a dismissal hearing notice. What must we do and by when?"

You have 1 week for ordinary dismissals, 3 days for immediate ones. Check that the employer provided complete social data and a reason. Objecting gives the employee a right to continued employment while the dispute is pending.

> "The employer wants to introduce an AI writing assistant. Do we have co-determination rights?"

Yes, under § 87 Abs. 1 Nr. 6 BetrVG if the tool can monitor output or behaviour — even indirectly. You should negotiate a Betriebsvereinbarung covering purpose, data access, and retention before rollout begins.

> "We're facing a mass layoff of 20 people. What's our role?"

The employer must consult you at least 30 days before filing the mass layoff notice with the employment agency. You can negotiate a Sozialplan — and force one through the Einigungsstelle if the employer refuses.

> "Can we force a Sozialplan if the employer refuses to negotiate?"

Yes. Under § 112 Abs. 4 BetrVG, the Sozialplan is erzwingbar — you can call the Einigungsstelle, which will set binding terms if negotiations fail.

---

## How to get it

Tell Claude to install it:

> Install this: https://github.com/googlarz/betriebsrat

Claude will handle the rest. Then just describe your situation and ask away.

---

## What it knows

- Every paragraph of the *Betriebsverfassungsgesetz* (BetrVG) — the German works constitution law
- Co-determination rights: when the works council can block, must consent, or is only informed
- Legal deadlines — including the critical **3-week window** to challenge a dismissal in court
- Sozialplan and severance calculations (Munich formula)
- AGG (discrimination), Mutterschutz (maternity), Elternzeit (parental leave), SGB IX (disability)
- Procedure violations — what happens if the employer didn't follow the rules

---

## License

Apache 2.0 — see [LICENSE](LICENSE).

*This tool provides legal orientation, not legal advice. For your specific situation, consult a labour law specialist or your trade union.*

---

<details>
<summary>For developers — CLI reference</summary>

### Install (CLI)

```bash
go install github.com/googlarz/betriebsrat/cmd/betriebsrat@latest
betriebsrat doctor
```

### Key commands

```bash
betriebsrat ask "Ich wurde entlassen. Was nun?"
betriebsrat rights-check "KI-Überwachungssystem einführen" --agent
betriebsrat decide "Massenentlassung 20 Personen" --json
betriebsrat sozialplan-calc --salary 3500 --years 7
betriebsrat nachteilsausgleich --salary 3500 --years 7 --no-ia-attempted
betriebsrat check-anhoerung "<text des Anhörungsschreibens>"
betriebsrat deadline kuendigung --from 2026-05-10
betriebsrat law 102
betriebsrat bv-template homeoffice
betriebsrat ki-check "Teams-Analysefunktion zur Produktivitätsmessung"
betriebsrat massenentlassung --employees 120 --affected 25
```

### MCP server

```bash
# Claude Code
claude mcp add betriebsrat betriebsrat-pp-mcp

# Claude Desktop
{
  "mcpServers": {
    "betriebsrat": { "command": "betriebsrat-pp-mcp" }
  }
}
```

### Output flags

`--json` · `--agent` · `--lang en` · `--lang de`

</details>

---

---

<a name="deutsch"></a>

# betriebsrat — Deutsche Dokumentation

**Arbeitsrechte kennen — für Arbeitnehmer und Betriebsratsmitglieder.**

Das deutsche Arbeitsrecht steckt voller Rechte, von denen die meisten Menschen nichts wissen. Dieses Tool macht sie sofort zugänglich — in einfacher Sprache, auf Deutsch oder Englisch, ohne dass Sie selbst Jura studiert haben müssen.

---

## Was kann ich fragen?

Beschreiben Sie einfach Ihre Situation. Hier sind Beispiele dafür, was das Tool bei der Antwort berücksichtigt:

**Als Arbeitnehmer:**

> „Habe ich Anspruch auf eine Abfindung? Ich arbeite hier seit 7 Jahren und verdiene 3.500 € monatlich."

Nach der Münchner Formel ergibt sich bei 7 Jahren und 3.500 € ein Richtwert von ca. 12.250 €. Ob ein Anspruch besteht, hängt davon ab, ob ein Sozialplan vereinbart wurde — oder übergangen wurde.

> „Ich habe eine Kündigung erhalten. Wurde der Betriebsrat ordnungsgemäß angehört?"

Vor jeder Kündigung muss der Betriebsrat angehört werden (§ 102 BetrVG). Fehlt die Anhörung oder waren die Angaben unvollständig, ist die Kündigung unwirksam — unabhängig vom angegebenen Grund.

> „Ich bin schwanger und mein Arbeitgeber will meinen Vertrag beenden."

Eine Kündigung während der Schwangerschaft ist nach dem Mutterschutzgesetz verboten — auch wenn der Arbeitgeber davon noch nichts weiß. Sie haben nach Erhalt der Kündigung 2 Wochen Zeit, ihn zu informieren; der Schutz gilt rückwirkend.

> „Mein Arbeitgeber hat eine Überwachungssoftware eingeführt. Musste der Betriebsrat zustimmen?"

Ja — jede Software, die Leistung oder Verhalten erfassen kann, bedarf der Zustimmung des Betriebsrats (§ 87 BetrVG). Wurde sie ohne Einigung eingeführt, kann der BR verlangen, dass der Einsatz gestoppt wird.

> „Ich wurde ohne mein Einverständnis versetzt. Ist das erlaubt?"

Versetzungen bedürfen der Zustimmung des Betriebsrats (§ 99 BetrVG). Wurde dieser nicht beteiligt oder ändern sich die Arbeitsbedingungen wesentlich, kann die Versetzung angefochten werden.

---

**Als Betriebsratsmitglied:**

> „Wir haben ein Anhörungsschreiben für eine Kündigung erhalten. Was müssen wir tun und bis wann?"

Bei ordentlichen Kündigungen haben Sie 1 Woche, bei fristlosen 3 Tage. Prüfen Sie, ob Sozialdaten und Kündigungsgrund vollständig angegeben wurden. Ein Widerspruch gibt dem Arbeitnehmer das Recht auf Weiterbeschäftigung während eines laufenden Verfahrens.

> „Der Arbeitgeber will einen KI-Schreibassistenten einführen. Haben wir ein Mitbestimmungsrecht?"

Ja, nach § 87 Abs. 1 Nr. 6 BetrVG, wenn das Tool Leistung oder Verhalten — auch mittelbar — erfassen kann. Handeln Sie eine Betriebsvereinbarung zu Zweck, Datenzugriff und Löschfristen aus, bevor die Einführung beginnt.

> „Uns droht eine Massenentlassung von 20 Personen. Was ist unsere Rolle?"

Der Arbeitgeber muss Sie mindestens 30 Tage vor Einreichung der Massenentlassungsanzeige bei der Agentur für Arbeit konsultieren. Sie können einen Sozialplan verhandeln — und ihn über die Einigungsstelle erzwingen, wenn der Arbeitgeber die Verhandlung verweigert.

> „Können wir einen Sozialplan erzwingen, wenn der Arbeitgeber nicht verhandeln will?"

Ja. Nach § 112 Abs. 4 BetrVG ist der Sozialplan erzwingbar — Sie können die Einigungsstelle anrufen, die bei gescheiterten Verhandlungen verbindliche Bedingungen festlegt.

---

## Wie bekomme ich das Tool?

Sagen Sie Claude, es soll installieren:

> Install this: https://github.com/googlarz/betriebsrat

Claude übernimmt den Rest. Dann beschreiben Sie einfach Ihre Situation.

---

## Was das Tool weiß

- Alle Paragrafen des Betriebsverfassungsgesetzes (BetrVG)
- Mitbestimmungsrechte: wann der Betriebsrat blockieren kann, zustimmen muss oder nur informiert wird
- Gesetzliche Fristen — einschließlich der kritischen **3-Wochen-Frist** für die Kündigungsschutzklage
- Sozialplan- und Abfindungsberechnungen (Münchner Formel)
- AGG (Diskriminierungsschutz), Mutterschutz, Elternzeit (BEEG), SGB IX (Schwerbehinderung)
- Verfahrensfehler — was passiert, wenn der Arbeitgeber die Regeln nicht eingehalten hat

---

## Lizenz

Apache 2.0 — siehe [LICENSE](LICENSE).

*Dieses Tool bietet rechtliche Orientierung, keine Rechtsberatung. Für Ihren konkreten Fall wenden Sie sich an einen Fachanwalt für Arbeitsrecht oder Ihre Gewerkschaft.*

---

<details>
<summary>Für Entwickler — CLI-Referenz</summary>

### Installation (CLI)

```bash
go install github.com/googlarz/betriebsrat/cmd/betriebsrat@latest
betriebsrat doctor
```

### Wichtige Befehle

```bash
betriebsrat ask "Ich wurde entlassen. Was nun?"
betriebsrat rights-check "KI-Überwachungssystem einführen" --agent
betriebsrat decide "Massenentlassung 20 Personen" --json
betriebsrat sozialplan-calc --salary 3500 --years 7
betriebsrat nachteilsausgleich --salary 3500 --years 7 --no-ia-attempted
betriebsrat check-anhoerung "<text des Anhörungsschreibens>"
betriebsrat deadline kuendigung --from 2026-05-10
betriebsrat law 102
betriebsrat bv-template homeoffice
betriebsrat ki-check "Teams-Analysefunktion zur Produktivitätsmessung"
betriebsrat massenentlassung --employees 120 --affected 25
```

### Ausgabeformate

`--json` · `--agent` · `--lang en` · `--lang de`

</details>
