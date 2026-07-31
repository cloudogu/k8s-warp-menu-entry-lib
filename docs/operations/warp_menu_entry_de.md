# WarpMenuEntry verwenden

Dogu-Entwickler registrieren ihre Anwendungen im Warp-Menü, indem sie eine
`WarpMenuEntry`-Custom-Resource im zugehörigen Namespace erstellen.

```yaml
apiVersion: k8s.cloudogu.com/v1
kind: WarpMenuEntry
metadata:
  name: my-dogu
  namespace: ecosystem
spec:
  displayName:
    de: "Mein Dogu"
    en: "My Dogu"
  category: "Development Apps"
  path: /my-dogu
```

## Spec-Felder

| Feld | Pflicht | Einschränkungen | Beschreibung |
|---|---|---|---|
| `displayName.de` | ja | 1–50 Zeichen | Deutscher Anzeigename im Menü. |
| `displayName.en` | ja | 1–50 Zeichen | Englischer Anzeigename im Menü. |
| `category` | ja | 1–50 Zeichen | Kategorie. Einen vordefinierten Schlüssel aus der `values.yaml` verwenden oder einen neuen angeben (Order: 9999). |
| `path` | ja | beginnt mit `/` | Serverrelativer URL-Pfad, z. B. `/my-dogu`. Darf keine Domain oder Schema enthalten. |
| `disabled` | nein | boolean | Bei `true` wird der Eintrag aus dem Menü ausgeblendet, ohne die Resource zu löschen. Standard: `false`. |

## Status-Conditions

Der Operator setzt nach jeder Reconciliation zwei Conditions am WarpMenuEntry.

| Condition | Status | Bedeutung |
|---|---|---|
| `Ready` | `True` | Eintrag ist gültig und das Menü wurde erfolgreich neu erstellt. |
| `Ready` | `False` | Eintrag hat einen Validierungsfehler (z. B. ungültiger Pfad, leere Kategorie) oder ein interner Fehler ist aufgetreten. Condition-Message oder Operator-Events prüfen. |
| `Visible` | `True` | Eintrag ist aktuell im Menü gerendert. |
| `Visible` | `False` | Eintrag ist ausgeblendet — entweder weil `disabled: true` gesetzt ist oder weil der Eintrag die Validierung nicht bestanden hat. |

Ressource inspizieren:

```shell
kubectl get warp -n ecosystem
kubectl describe warp my-dogu -n ecosystem
```

## Eintrag vorübergehend ausblenden

`disabled: true` setzen, um einen Eintrag zu verstecken, ohne die Resource zu löschen:

```yaml
spec:
  disabled: true
```
