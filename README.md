# YAML Toolbox

A collection of small YAML tools that can be used from the CLI or a GitHub Action

### Splitter

Given a Kubernetes multi-doc YAML file, split the file into separate files.

The following GitHub Action workflow will split the contents of api.yaml into separate documents, in the api directory. The filenames will be [metadata.name]-[kind].yaml.

```yaml
      - name: Split the api doc
        uses: marccampbell/yaml-toolbox/action/split@master
        with:
          path: ./api.yaml
          out: ./api
```

### Remarshaler

Given a YAML document or a directory of YAML documents, unmarshal and remarshal to create deterministic YAML. This reorders the YAML document to be consistent, and thereby creates smaller diffs. It also (importantly) doesn't wrap lines at 80 chars, which is useful when using a template functions in YAML for tools such as [KOTS](https://kots.io) or [Helm](https://helm.sh).

The following GitHub Action workflow will remarshal all of the documetns in the kots directory:

```yaml
      - name: Remarshal the release
        uses: marccampbell/yaml-toolbox/action/remarshal@master
        with:
          path: ./kots
```