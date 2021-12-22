# DVF Converter

### Datasets
- Raw Data : https://www.data.gouv.fr/fr/datasets/5c4ae55a634f4117716d5656/#resources (TXT format)
- Raw Geo Data : https://www.data.gouv.fr/fr/datasets/5cc1b94a634f4165e96436c1/#resources (CSV format)
### Features
**TXT to CSV**
```bash
$ dvf-converter TxtToCsv  -i ./snippets/base_data/snippet-2021-S1.txt -o voila -d "|" 
```

**CSV to JSON**
`WIP`

**CSV to JSON**
`WIP`

#### Todos :
- Improve errors management
- Add tests
- Add TXT to JSON feat
- Add CSV to JSON feat
- Improve delimiters management 
- Add documentation