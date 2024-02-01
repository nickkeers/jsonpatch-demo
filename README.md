# JSON Patch Demo

This is a small demo of how [JSONPatch](https://jsonpatch.com) works, in this code I demo how you could store JSON patch objects
as "changes" in a JSON document you can store in a document store like MongoDB. Once you have your original document and changes
you can easily "materialise" the full document from the patches which will let you have a way to do changelogs / audit logging
of changes to your documents.

To run:

```bash 
go run main.go 
```

Change [patch.json](patch.json) as you see fit, following the [JSONPatch operations](https://jsonpatch.com/#operations) to
see the effect on the final document.