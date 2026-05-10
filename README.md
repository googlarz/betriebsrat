# betriebsrat

**Know your rights at work. For employees and works council members.**

> **[🇩🇪 Deutsche Version weiter unten](#deutsch)**

German labour law is detailed, deadline-driven, and full of rights that most people don't know they have. This tool makes that knowledge instantly accessible — whether you're a works council (*Betriebsrat*) member trying to do your job, or an employee trying to understand what your employer can and can't do to you.

No API key. No login. No subscription. Runs entirely offline from an embedded BetrVG knowledge base.

---

## Who is this for?

### Employees (Arbeitnehmer)

You're facing a dismissal, transfer, restructuring, or a new monitoring system at work. You want to know: **was the procedure followed? What can I claim? What are my rights?**

| Your question | Command |
|---|---|
| Can my employer fire me without asking the BR? | `consequences kündigung` |
| My company is restructuring — what am I entitled to? | `sozialplan-calc` + `nachteilsausgleich` |
| They're introducing monitoring software — can the BR stop it? | `rights-check` + `ki-check` |
| Is my dismissal valid if the BR wasn't properly consulted? | `check-anhoerung` + `consequences kündigung` |
| I'm being transferred — did they need BR approval? | `rights-check` + `consequences versetzung` |
| What's the BR even allowed to do in my situation? | `decide` |
| How much severance am I entitled to? | `sozialplan-calc` |
| The company skipped the Interessenausgleich — can I claim more? | `nachteilsausgleich` |
| Is there a works council at companies my size? | `law 1` |

### Works Council Members (Betriebsratsmitglieder)

You're managing a co-determination situation — dismissal hearings, a restructuring, a new IT system, Betriebsvereinbarung negotiations. You need the right answer fast, the correct deadline, and the right document.

| Your question | Command |
|---|---|
| Do we have co-determination rights here? | `rights-check` |
| What kind of right is it — can we actually block this? | `codetermination-type` |
| What's our deadline and what happens if we miss it? | `deadline` + `consequences` |
| What should we do, step by step? | `decide` |
| Is this Anhörungsschreiben valid? Does our clock run? | `check-anhoerung` |
| Draft a Betriebsvereinbarung for this | `bv-template` |
| Does the Tarifvertrag block us from writing a BV? | `tarifvertrag-check` |
| Does this AI system trigger § 87 Nr. 6? | `ki-check` |
| We need to request information from the employer | `auskunft` |
| What are our strongest Widerspruch grounds? | `widerspruch-check` |
| Calculate Sozialplan for all affected employees | `sozialplan-calc --csv` |
| Request training time off under § 37 Abs. 6 | `schulungsantrag` |
| Do these redundancies trigger § 17 KSchG? | `massenentlassung` |

---

## Install

### Recommended (CLI + Claude Code skill)

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

### CLI only

```bash
npx -y @mvanhorn/printing-press install betriebsrat --cli-only
```

### From source (Go 1.21+)

```bash
git clone https://github.com/googlarz/betriebsrat.git
cd betriebsrat
go install github.com/googlarz/betriebsrat/cmd/betriebsrat@latest
```

### Pre-built binary

Download from [Releases](../../releases). On macOS: `xattr -d com.apple.quarantine betriebsrat`. On Linux: `chmod +x betriebsrat`.

### Verify

```bash
betriebsrat doctor
```

---

## Quick Start

```bash
# Populate local knowledge base (run once, everything else works offline)
betriebsrat sync

# --- AS AN EMPLOYEE ---

# "My employer is dismissing me — was the BR properly consulted?"
betriebsrat consequences kündigung --lang en

# "I'm in a restructuring — how much am I entitled to?"
betriebsrat sozialplan-calc --salary 4500 --years 10 --age 45

# "The employer didn't attempt an Interessenausgleich — can I claim more?"
betriebsrat nachteilsausgleich --salary 4500 --years 10 --age 45 --measure "Standortschließung" --no-ia-attempted

# "New monitoring system at work — does the BR have to agree?"
betriebsrat rights-check "employer introducing performance monitoring software" --lang en

# --- AS A BR MEMBER ---

# "Do we have co-determination rights here?"
betriebsrat rights-check "Arbeitgeber will Überwachungssoftware einführen"

# "What's our deadline?"
betriebsrat deadline "ordentliche Kündigung" --from $(date +%Y-%m-%d)

# "Full decision support for a situation"
betriebsrat decide "Betrieb soll nach München verlagert werden"

# "Draft a BV"
betriebsrat bv-template homeoffice --employer "Musterfirma GmbH"

# English output on any command
betriebsrat rights-check "mass layoff of 20 employees" --lang en
```

---

## Commands

### Understanding rights and consequences

#### `rights-check` — Does the BR have co-determination rights in this situation?

Useful for **employees** ("should the BR have been involved?") and **BR members** ("do we have a say here?").

```bash
betriebsrat rights-check "Arbeitgeber will Überwachungssoftware einführen"
betriebsrat rights-check "employer wants to transfer me to another city" --lang en
betriebsrat rights-check "Homeoffice-Regelung" --json
```

Output: applicable §§, co-determination type (erzwingbar / Mitwirkung / Unterrichtung / keine), recommendation.

#### `codetermination-type` — What kind of right is it?

Classifies as: **Mitbestimmung (erzwingbar)** — employer can be blocked · Zustimmungsvorbehalt · Mitwirkung · Beratung · Unterrichtung · keine

```bash
betriebsrat codetermination-type "Versetzung"
betriebsrat codetermination-type "Einführung Schichtplan"
```

#### `decide` — Full decision support

For **employees**: understand the full picture of a situation — which rights apply, what can happen next.  
For **BR members**: classified situation, applicable rights, prioritised action plan, deadlines.

```bash
betriebsrat decide "Betrieb soll nach München verlagert werden"
betriebsrat decide "mass layoff of 20 employees" --lang en --agent
```

#### `consequences` — What happens if procedure is violated?

For **employees**: if the BR wasn't consulted, or deadlines were missed — what does that mean for the validity of the measure against you?  
For **BR members**: what leverage do you have, and what happens if *you* miss the window?

```bash
betriebsrat consequences kündigung     # dismissal without proper hearing
betriebsrat consequences einstellung   # hire without BR consent
betriebsrat consequences versetzung    # transfer without BR consent
betriebsrat consequences betriebsänderung  # restructuring without Interessenausgleich
betriebsrat consequences software      # monitoring system without BV
betriebsrat consequences br-deadline   # BR didn't respond in time
betriebsrat consequences kündigung --lang en
```

#### `checklist` — Step-by-step action checklist

```bash
betriebsrat checklist "Betriebsänderung"
betriebsrat checklist "Massenentlassung" --lang en
```

---

### Deadlines

#### `deadline` — Calculate legal response deadlines

For **employees**: know when the BR's window closes — after that, their silence counts as consent.  
For **BR members**: calculate the exact date you must respond by.

```bash
betriebsrat deadline "ordentliche Kündigung" --from 2026-05-10
betriebsrat deadline "außerordentliche Kündigung" --from 2026-05-10
# Export to calendar
betriebsrat deadline "ordentliche Kündigung" --from 2026-05-10 --ical > frist.ics
```

Built-in: § 102 (7 days ordinary, 3 days extraordinary dismissal), § 99 hiring/transfer (1 week), § 17 KSchG mass dismissal (30 days), and more.

---

### Calculations

#### `sozialplan-calc` — How much severance is an employee entitled to?

For **employees**: find out what you're owed under the Sozialplan. For **BR members**: calculate and compare entitlements for all affected staff.

```bash
betriebsrat sozialplan-calc --salary 4500 --years 12 --age 48
betriebsrat sozialplan-calc --salary 6000 --years 20 --age 55 --disabled --factor 0.8
# Batch for all affected employees (CSV)
betriebsrat sozialplan-calc --csv employees.csv --agent
```

CSV format: `name,salary,years,age,disabled,children[,factor[,max_cap]]`

#### `nachteilsausgleich` — Can I claim more if the employer skipped the Interessenausgleich?

For **employees**: when the employer implemented a restructuring without attempting an Interessenausgleich with the BR, each affected employee has an individual compensation claim under § 113 BetrVG.

```bash
betriebsrat nachteilsausgleich --salary 4500 --years 8 --age 42 --measure "Standortschließung" --no-ia-attempted
betriebsrat nachteilsausgleich --salary 6000 --years 15 --age 55 --measure "relocation abroad" --ia-deviated --lang en
```

---

### Legal checks

#### `check-anhoerung` — Is the dismissal hearing notice valid? Does the deadline run?

For **employees**: an incomplete Anhörungsschreiben means the dismissal may be void, and the BR's deadline has not yet started.  
For **BR members**: verify completeness before the clock runs out.

```bash
betriebsrat check-anhoerung "Sehr geehrter Betriebsrat, wir beabsichtigen..."
betriebsrat check-anhoerung --file anhörung.txt --type ordentlich
```

#### `massenentlassung` — Do these redundancies trigger § 17 KSchG?

```bash
betriebsrat massenentlassung --employees 120 --planned-dismissals 15
betriebsrat massenentlassung --employees 50 --planned-dismissals 8 --lang en
```

If the threshold is met and the employer skipped the mandatory BA notification procedure, the dismissals are void.

#### `widerspruch-check` — What are the strongest Widerspruch grounds?

For **employees**: a BR Widerspruch means you have a right to continued employment during your court case.  
For **BR members**: rank grounds by legal strength and get the draft text.

```bash
betriebsrat widerspruch-check --reason "Sozialauswahl fehlerhaft" --age 52 --years 18
betriebsrat widerspruch-check --reason "freier Arbeitsplatz vorhanden" --lang en
```

#### `ki-check` — Does this AI/IT system trigger § 87 Nr. 6 co-determination?

```bash
betriebsrat ki-check --system "Leistungsmonitoring-Dashboard" --monitors-performance --influences-hr
betriebsrat ki-check --system "AI recruitment screener" --auto-decision --lang en
```

#### `tarifvertrag-check` — Does the Tarifvertrag block a planned BV?

Always run before drafting a BV in a TV-regulated area.

```bash
betriebsrat tarifvertrag-check --topic lohn --tv-type "Branchentarifvertrag" --tv-covers
betriebsrat tarifvertrag-check --topic homeoffice --no-tv-covers
betriebsrat tarifvertrag-check --topic software --lang en
```

---

### Document generation

#### `bv-template` — Betriebsvereinbarung skeleton

```bash
betriebsrat bv-template homeoffice --employer "Musterfirma GmbH"
betriebsrat bv-template software --agent
betriebsrat bv-template videoüberwachung --employer "Firma AG"
betriebsrat bv-template leistungsbeurteilung --employer "TechCo GmbH"
```

Topics: `homeoffice` · `software` · `arbeitszeit` · `datenschutz` · `videoüberwachung` · `leistungsbeurteilung`

#### `auskunft` — § 80 BetrVG information request to the employer

```bash
betriebsrat auskunft --topic sozialdaten --reason "Prüfung Sozialauswahl § 102" --employer "Firma GmbH"
betriebsrat auskunft --topic ki --reason "Einführung KI-Bewertungssystem" --deadline-days 10
betriebsrat auskunft --topic custom --custom "Überstundenaufstellungen der letzten 12 Monate"
```

Topics: `sozialdaten` · `stellenplan` · `gehaelter` · `planung` · `auswahlrichtlinien` · `ki` · `wirtschaft` · `custom`

#### `letter` — Formal BR letters

```bash
betriebsrat letter kündigung --type widerspruch
betriebsrat letter betriebsänderung --type einigungsstelle
```

#### `schulungsantrag` — § 37 Abs. 6 training request

```bash
betriebsrat schulungsantrag --topic betrvg --employer "Musterfirma GmbH"
betriebsrat schulungsantrag --topic kuendigung --provider "ver.di Bildung" --employer "AG GmbH"
```

Topics: `betrvg` · `arbeitsrecht` · `betriebsrat-praxis` · `kuendigung` · `sozialplan` · `datenschutz` · `gesundheit` · `custom`

#### `protokoll` — BR meeting minutes template

```bash
betriebsrat protokoll --topic "Anhörung Kündigung Müller" --br-size 9 --date 2026-05-10
```

---

### Legal reference

#### `law` — BetrVG paragraph in plain language

```bash
betriebsrat law 87    # co-determination on monitoring systems
betriebsrat law 102   # dismissal hearing
betriebsrat law 113   # Nachteilsausgleich individual claims
betriebsrat law kündigung   # keyword search
```

#### `prepare-meeting` — Meeting preparation

```bash
betriebsrat prepare-meeting "Einführung KI-System"
betriebsrat prepare-meeting "Betriebsänderung" --lang en
```

#### `context` — Store company profile for calibrated advice

```bash
betriebsrat context set --employees 85 --tariff true --br-size 7 --sector "Handel"
betriebsrat context show
```

---

## Output formats

```bash
# Human-readable (default)
betriebsrat rights-check "Überwachungssoftware"

# JSON — for scripting, piping, AI agents
betriebsrat rights-check "Überwachungssoftware" --json

# Agent mode — JSON + compact + no prompts in one flag
betriebsrat rights-check "Überwachungssoftware" --agent

# English output
betriebsrat rights-check "Überwachungssoftware" --lang en

# Filter specific fields
betriebsrat rights-check "Überwachungssoftware" --json --select summary,recommendation
```

Exit codes: `0` success · `2` usage error · `3` not found · `5` API error · `7` rate limited · `10` config error

---

## Use with Claude Code

Install the agent skill:

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

Describe your situation in plain language — employee or BR member, German or English. The skill detects who you are, runs the right commands, and chains follow-ups automatically (e.g. incomplete Anhörung → consequences; restructuring → sozialplan-calc + nachteilsausgleich).

## Use as an MCP server

```bash
# Claude Code
claude mcp add betriebsrat betriebsrat-pp-mcp

# Claude Desktop — add to claude_desktop_config.json
{
  "mcpServers": {
    "betriebsrat": { "command": "betriebsrat-pp-mcp" }
  }
}
```

---

## Offline capability

All decision-support commands run offline after the initial `sync`:

`rights-check` · `decide` · `deadline` · `checklist` · `law` · `codetermination-type` · `consequences` · `letter` · `sozialplan-calc` · `nachteilsausgleich` · `massenentlassung` · `widerspruch-check` · `protokoll` · `auskunft` · `ki-check` · `schulungsantrag` · `tarifvertrag-check` · `bv-template` · `context` · `check-anhoerung`

---

## License

Apache 2.0 — see [LICENSE](LICENSE).

*Legal notice: This tool is for orientation purposes and does not replace professional legal advice. For complex individual cases, consult a labour law specialist or your trade union.*

---

---

<a name="deutsch"></a>

# betriebsrat — Deutsche Dokumentation

**Arbeitsrechte verstehen. Für Arbeitnehmer und Betriebsratsmitglieder.**

Deutsches Arbeitsrecht ist detailliert, fristengebunden und voller Rechte, von denen die meisten Menschen nichts wissen. Dieses Tool macht dieses Wissen sofort zugänglich — egal ob Sie Betriebsratsmitglied sind und Ihre Arbeit erledigen wollen, oder Arbeitnehmer, der verstehen möchte, was der Arbeitgeber mit Ihnen machen kann und was nicht.

Kein API-Key. Kein Login. Kein Abo. Läuft vollständig offline aus einer eingebetteten BetrVG-Wissensdatenbank.

---

## Für wen ist das?

### Arbeitnehmer

Sie stehen vor einer Kündigung, Versetzung, Umstrukturierung oder einer neuen Überwachungstechnik am Arbeitsplatz. Sie wollen wissen: **Wurde das Verfahren eingehalten? Was steht mir zu? Was sind meine Rechte?**

| Ihre Frage | Befehl |
|---|---|
| Darf mein Arbeitgeber mich kündigen, ohne den BR zu fragen? | `consequences kündigung` |
| Mein Unternehmen restrukturiert — was steht mir zu? | `sozialplan-calc` + `nachteilsausgleich` |
| Sie führen Überwachungssoftware ein — kann der BR das stoppen? | `rights-check` + `ki-check` |
| Ist meine Kündigung wirksam, wenn der BR nicht richtig angehört wurde? | `check-anhoerung` + `consequences kündigung` |
| Ich werde versetzt — brauchten sie die Zustimmung des BR? | `rights-check` + `consequences versetzung` |
| Was darf der BR überhaupt in meiner Situation tun? | `decide` |
| Wie viel Abfindung steht mir zu? | `sozialplan-calc` |
| Der AG hat keinen Interessenausgleich versucht — kann ich mehr verlangen? | `nachteilsausgleich` |

### Betriebsratsmitglieder

Sie bearbeiten eine Mitbestimmungssituation — Kündigungsanhörungen, Umstrukturierungen, neue IT-Systeme, BV-Verhandlungen. Sie brauchen die richtige Antwort schnell, die korrekte Frist und das richtige Dokument.

| Ihre Frage | Befehl |
|---|---|
| Haben wir hier ein Mitbestimmungsrecht? | `rights-check` |
| Welche Art von Recht ist es — können wir das wirklich blockieren? | `codetermination-type` |
| Was ist unsere Frist und was passiert, wenn wir sie versäumen? | `deadline` + `consequences` |
| Was sollen wir jetzt tun, Schritt für Schritt? | `decide` |
| Ist dieses Anhörungsschreiben vollständig? Läuft unsere Frist? | `check-anhoerung` |
| Betriebsvereinbarung für dieses Thema erstellen | `bv-template` |
| Sperrt der Tarifvertrag uns, eine BV zu schreiben? | `tarifvertrag-check` |
| Löst dieses KI-System § 87 Nr. 6 aus? | `ki-check` |
| Wir müssen Informationen vom Arbeitgeber anfordern | `auskunft` |
| Was sind unsere stärksten Widerspruchsgründe? | `widerspruch-check` |
| Sozialplan für alle betroffenen Mitarbeiter berechnen | `sozialplan-calc --csv` |
| Schulungsantrag nach § 37 Abs. 6 stellen | `schulungsantrag` |
| Lösen diese Entlassungen § 17 KSchG aus? | `massenentlassung` |

---

## Installation

### Empfohlen (CLI + Claude Code Skill)

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

### Nur CLI

```bash
npx -y @mvanhorn/printing-press install betriebsrat --cli-only
```

### Aus dem Quellcode (Go 1.21+)

```bash
git clone https://github.com/googlarz/betriebsrat.git
cd betriebsrat
go install github.com/googlarz/betriebsrat/cmd/betriebsrat@latest
```

### Installation prüfen

```bash
betriebsrat doctor
```

---

## Schnellstart

```bash
# Wissensdatenbank befüllen (einmalig, danach alles offline)
betriebsrat sync

# --- ALS ARBEITNEHMER ---

# "Mein AG kündigt mir — wurde der BR richtig angehört?"
betriebsrat consequences kündigung

# "Ich bin von einer Umstrukturierung betroffen — was steht mir zu?"
betriebsrat sozialplan-calc --salary 4500 --years 10 --age 45

# "Der AG hat keinen IA versucht — kann ich mehr verlangen?"
betriebsrat nachteilsausgleich --salary 4500 --years 10 --age 45 --measure "Standortschließung" --no-ia-attempted

# "Neue Überwachungssoftware — muss der BR zustimmen?"
betriebsrat rights-check "Arbeitgeber führt Leistungsmonitoring ein"

# --- ALS BETRIEBSRATSMITGLIED ---

# "Haben wir ein Mitbestimmungsrecht?"
betriebsrat rights-check "Arbeitgeber will Überwachungssoftware einführen"

# "Was ist unsere Frist?"
betriebsrat deadline "ordentliche Kündigung" --from $(date +%Y-%m-%d)

# "Vollständige Entscheidungshilfe"
betriebsrat decide "Betrieb soll nach München verlagert werden"

# "Englische Ausgabe bei jedem Befehl"
betriebsrat rights-check "Massenentlassung 15 Mitarbeiter" --lang en
```

---

## Befehle im Detail

### Rechte und Konsequenzen verstehen

#### `rights-check` — Hat der BR ein Mitbestimmungsrecht in dieser Situation?

Nützlich für **Arbeitnehmer** ("Hätte der BR hier einbezogen werden müssen?") und **BR-Mitglieder** ("Haben wir hier ein Wort mitzureden?").

```bash
betriebsrat rights-check "Arbeitgeber will Überwachungssoftware einführen"
betriebsrat rights-check "Ich soll in eine andere Stadt versetzt werden"
```

#### `consequences` — Was passiert bei Verfahrensfehlern?

Für **Arbeitnehmer**: Wenn der BR nicht angehört wurde oder Fristen versäumt wurden — was bedeutet das für die Wirksamkeit der Maßnahme gegen Sie?  
Für **BR-Mitglieder**: Welchen Hebel haben Sie, und was passiert, wenn *Sie* die Frist versäumen?

```bash
betriebsrat consequences kündigung      # Kündigung ohne ordentliche Anhörung
betriebsrat consequences betriebsänderung  # Umstrukturierung ohne Interessenausgleich
betriebsrat consequences software       # Überwachungssystem ohne BV
betriebsrat consequences br-deadline    # BR hat nicht rechtzeitig geantwortet
```

#### `decide` — Vollständige Situationsanalyse

```bash
betriebsrat decide "Betrieb soll nach München verlagert werden"
betriebsrat decide "Arbeitgeber kündigt 15 Mitarbeiter"
```

---

### Fristen

#### `deadline` — Gesetzliche Fristen berechnen

Für **Arbeitnehmer**: Wissen wann das BR-Zeitfenster schließt — danach gilt Schweigen als Zustimmung.  
Für **BR-Mitglieder**: Das genaue Datum berechnen, bis zu dem Sie antworten müssen.

```bash
betriebsrat deadline "ordentliche Kündigung" --from 2026-05-10
betriebsrat deadline "außerordentliche Kündigung" --from 2026-05-10
# Als Kalenderdatei exportieren
betriebsrat deadline "ordentliche Kündigung" --from 2026-05-10 --ical > frist.ics
```

---

### Berechnungen

#### `sozialplan-calc` — Wie viel Abfindung steht einem Mitarbeiter zu?

Für **Arbeitnehmer**: Herausfinden, was Ihnen zusteht.  
Für **BR-Mitglieder**: Ansprüche aller betroffenen Mitarbeiter berechnen und vergleichen.

```bash
betriebsrat sozialplan-calc --salary 4500 --years 12 --age 48
betriebsrat sozialplan-calc --salary 6000 --years 20 --age 55 --disabled
# Stapelmodus für alle betroffenen Mitarbeiter (CSV)
betriebsrat sozialplan-calc --csv mitarbeiter.csv --agent
```

#### `nachteilsausgleich` — Kann ich mehr verlangen, weil der AG den Interessenausgleich übergangen hat?

Für **Arbeitnehmer**: Wenn der AG eine Betriebsänderung ohne Versuch eines Interessenausgleichs durchgeführt hat, hat jeder betroffene Arbeitnehmer einen individuellen Abfindungsanspruch nach § 113 BetrVG.

```bash
betriebsrat nachteilsausgleich --salary 4500 --years 8 --age 42 --measure "Standortschließung" --no-ia-attempted
betriebsrat nachteilsausgleich --salary 6000 --years 15 --age 55 --measure "Verlagerung ins Ausland" --ia-deviated
```

---

### Rechtliche Prüfungen

#### `check-anhoerung` — Ist das Anhörungsschreiben vollständig? Läuft die Frist?

Für **Arbeitnehmer**: Ein unvollständiges Anhörungsschreiben bedeutet, dass die Kündigung möglicherweise unwirksam ist und die BR-Frist noch gar nicht zu laufen begonnen hat.

```bash
betriebsrat check-anhoerung "Sehr geehrter Betriebsrat, wir beabsichtigen..."
betriebsrat check-anhoerung --file anhörung.txt --type ordentlich
```

#### `widerspruch-check` — Was sind die stärksten Widerspruchsgründe?

Für **Arbeitnehmer**: Ein BR-Widerspruch bedeutet, dass Sie das Recht auf Weiterbeschäftigung während Ihres Gerichtsverfahrens haben (§ 102 Abs. 5 BetrVG).

```bash
betriebsrat widerspruch-check --reason "Sozialauswahl fehlerhaft" --age 52 --years 18
betriebsrat widerspruch-check --reason "freier Arbeitsplatz vorhanden"
```

#### `massenentlassung` — Lösen diese Entlassungen § 17 KSchG aus?

```bash
betriebsrat massenentlassung --employees 120 --planned-dismissals 15
```

Wenn der Schwellenwert überschritten wird und der AG das Pflichtverfahren bei der BA übersprungen hat, sind die Kündigungen unwirksam.

#### `ki-check` + `tarifvertrag-check`

```bash
betriebsrat ki-check --system "KI-Recruiting-Tool" --auto-decision
betriebsrat tarifvertrag-check --topic lohn --tv-covers
```

---

### Dokumentenerstellung

```bash
# Betriebsvereinbarung
betriebsrat bv-template homeoffice --employer "Musterfirma GmbH"
betriebsrat bv-template videoüberwachung --employer "Firma AG"
betriebsrat bv-template leistungsbeurteilung --employer "TechCo GmbH"

# Auskunftsverlangen an den AG
betriebsrat auskunft --topic sozialdaten --employer "Firma GmbH"
betriebsrat auskunft --topic ki --reason "Einführung KI-Bewertungssystem"

# Schulungsantrag
betriebsrat schulungsantrag --topic betrvg --employer "Musterfirma GmbH"

# BR-Sitzungsprotokoll
betriebsrat protokoll --topic "Anhörung Kündigung Müller" --br-size 9
```

---

### Nachschlagewerk

```bash
# BetrVG-Paragraf in Alltagssprache
betriebsrat law 87    # Mitbestimmung bei Überwachungstechnik
betriebsrat law 102   # Kündigungsanhörung
betriebsrat law 113   # Nachteilsausgleich

# Stichwortsuche
betriebsrat law kündigung
betriebsrat law versetzung

# Betriebsprofil für kalibrierte Beratung
betriebsrat context set --employees 85 --tariff true --br-size 7
```

---

## Verwendung mit Claude Code

Nach der Installation beschreiben Sie Ihre Situation auf Deutsch oder Englisch — als Arbeitnehmer oder BR-Mitglied. Der Skill erkennt den Kontext, führt die richtigen Befehle aus und kettet Folgeschritte automatisch (z.B. unvollständige Anhörung → Konsequenzen; Umstrukturierung → Sozialplan + Nachteilsausgleich).

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

---

## Offline-Fähigkeit

Alle Entscheidungshilfe-Befehle laufen nach dem einmaligen `sync` vollständig offline:

`rights-check` · `decide` · `deadline` · `checklist` · `law` · `codetermination-type` · `consequences` · `letter` · `sozialplan-calc` · `nachteilsausgleich` · `massenentlassung` · `widerspruch-check` · `protokoll` · `auskunft` · `ki-check` · `schulungsantrag` · `tarifvertrag-check` · `bv-template` · `context` · `check-anhoerung`

---

## Lizenz

Apache 2.0 — siehe [LICENSE](LICENSE).

*Rechtlicher Hinweis: Dieses Tool dient der rechtlichen Orientierung und ersetzt keine anwaltliche Beratung. Bei komplexen Einzelfällen empfehlen wir die Konsultation eines Fachanwalts für Arbeitsrecht oder Ihrer Gewerkschaft.*
