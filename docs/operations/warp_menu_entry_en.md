## Using WarpMenuEntry 

Dogu developers register their applications in the warp menu by creating a `WarpMenuEntry`
custom resource in the operator namespace.

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

## Spec fields

| Field | Required | Constraints | Description |
|---|---|---|---|
| `displayName.de` | yes | 1–50 characters | German display name shown in the menu. |
| `displayName.en` | yes | 1–50 characters | English display name shown in the menu. |
| `category` | yes | 1–50 characters | Category identifier. Use a pre-defined key from `values.yaml` or provide any new identifier to create an ad-hoc category (order will be 9999). |
| `path` | yes | starts with `/` | Server-relative URL path, e.g. `/my-dogu`. Must not include a domain or scheme. |
| `disabled` | no | boolean | When `true`, the entry is hidden from the menu without deleting the resource. Defaults to `false`. |

## Status conditions

The operator sets two conditions on every WarpMenuEntry after reconciliation.

| Condition | Status | Meaning |
|---|---|---|
| `Ready` | `True` | Entry is valid and the menu was rebuilt successfully. |
| `Ready` | `False` | Entry has a validation error (e.g. bad path, empty category) or an internal error occurred. Check the condition message or operator events for details. |
| `Visible` | `True` | Entry is currently rendered in the menu. |
| `Visible` | `False` | Entry is hidden — either `disabled: true` is set or the entry failed validation. |

Inspect a resource with:

```shell
kubectl get warp -n ecosystem
kubectl describe warp my-dogu -n ecosystem
```

## Temporarily hiding an entry

Set `disabled: true` to remove an entry from the menu without deleting the resource:

```yaml
spec:
  disabled: true
```
