package storage

import (
	"bytes"
	"io"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg/errors"

	// "fmt"
	// "strings"
	// "log"
	// "google.golang.org/api/iterator"
)

// readFile reads the named file in Google Cloud Storage.
func ReadFile(fileName string) ([]byte, error) {

	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bucket := client.Bucket("atsur-registry-files")

	w := &bytes.Buffer{}

	io.WriteString(w, "\nAbbreviated file content (first line and last 1K):\n")

	rc, err := bucket.Object(fileName).NewReader(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "cannot read bucket")
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil read all")
	}

	return slurp, nil
}


// type demo struct {
// 	bucketName string

// 	cleanUp []string

// 	failed bool
// }

// func (d *demo) createFile(fileName string) error {

// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}
// 	fmt.Fprintf(w, "Creating file /%v/%v\n", d.bucketName, fileName)

// 	wc := bucket.Object(fileName).NewWriter(context.Background())
// 	wc.ContentType = "text/plain"
// 	wc.Metadata = map[string]string{
// 		"x-goog-meta-foo": "foo",
// 		"x-goog-meta-bar": "bar",
// 	}
// 	d.cleanUp = append(d.cleanUp, fileName)

// 	if _, err := wc.Write([]byte("abcde\n")); err != nil {
// 		return err
// 	}
// 	if _, err := wc.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
// 		return err
// 	}
// 	if err := wc.Close(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (d *demo) copyFile(fileName string) error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}
// 	copyName := fileName + "-copy"
// 	fmt.Fprintf(w, "Copying file /%v/%v to /%v/%v:\n", d.bucketName, fileName, d.bucketName, copyName)

// 	obj, err := bucket.Object(copyName).CopierFrom(bucket.Object(fileName)).Run(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	d.cleanUp = append(d.cleanUp, copyName)

// 	d.dumpStats(obj)
// 	return nil
// }

// func (d *demo) dumpStats(obj *storage.ObjectAttrs) {

// 	w := &bytes.Buffer{}
// 	fmt.Fprintf(w, "(filename: /%v/%v, ", obj.Bucket, obj.Name)
// 	fmt.Fprintf(w, "ContentType: %q, ", obj.ContentType)
// 	fmt.Fprintf(w, "ACL: %#v, ", obj.ACL)
// 	fmt.Fprintf(w, "Owner: %v, ", obj.Owner)
// 	fmt.Fprintf(w, "ContentEncoding: %q, ", obj.ContentEncoding)
// 	fmt.Fprintf(w, "Size: %v, ", obj.Size)
// 	fmt.Fprintf(w, "MD5: %q, ", obj.MD5)
// 	fmt.Fprintf(w, "CRC32C: %q, ", obj.CRC32C)
// 	fmt.Fprintf(w, "Metadata: %#v, ", obj.Metadata)
// 	fmt.Fprintf(w, "MediaLink: %q, ", obj.MediaLink)
// 	fmt.Fprintf(w, "StorageClass: %q, ", obj.StorageClass)
// 	if !obj.Deleted.IsZero() {
// 		fmt.Fprintf(w, "Deleted: %v, ", obj.Deleted)
// 	}
// 	fmt.Fprintf(w, "Updated: %v)\n", obj.Updated)
// }

// func (d *demo) statFile(fileName string) error {

// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}

// 	io.WriteString(w, "\nFile stat:\n")

// 	obj, err := bucket.Object(fileName).Attrs(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	d.dumpStats(obj)
// 	return nil
// }

// func (d *demo) createListFiles() {
// 	w := &bytes.Buffer{}
// 	io.WriteString(w, "\nCreating more files for listbucket...\n")
// 	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/"} {
// 		d.createFile(n)
// 	}
// }

// func (d *demo) listBucket() error {

// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}

// 	io.WriteString(w, "\nListbucket result:\n")

// 	query := &storage.Query{Prefix: "foo"}
// 	it := bucket.Objects(context.Background(), query)
// 	for {
// 		obj, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		d.dumpStats(obj)
// 	}
// 	return nil
// }

// func (d *demo) listDir(name, indent string) error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}

// 	query := &storage.Query{Prefix: name, Delimiter: "/"}
// 	it := bucket.Objects(context.Background(), query)
// 	for {
// 		obj, err := it.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		if obj.Prefix == "" {
// 			fmt.Fprint(w, indent)
// 			d.dumpStats(obj)
// 			continue
// 		}
// 		fmt.Fprintf(w, "%v(directory: /%v/%v)\n", indent, d.bucketName, obj.Prefix)
// 		d.listDir(obj.Prefix, indent+"  ")
// 	}
// 	return nil
// }

// func (d *demo) listBucketDirMode() {
// 	w := &bytes.Buffer{}
// 	io.WriteString(w, "\nListbucket directory mode result:\n")
// 	d.listDir("b", "")
// }

// func (d *demo) dumpDefaultACL() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}

// 	acl, err := bucket.ACL().List(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	for _, v := range acl {
// 		fmt.Fprintf(w, "Scope: %q, Permission: %q\n", v.Entity, v.Role)
// 	}
// 	return nil
// }

// func (d *demo) defaultACL() {
// 	w := &bytes.Buffer{}
// 	io.WriteString(w, "\nDefault object ACL:\n")
// 	d.dumpDefaultACL()
// }

// func (d *demo) putDefaultACLRule() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}
// 	io.WriteString(w, "\nPut Default object ACL Rule:\n")

// 	err = bucket.DefaultObjectACL().Set(context.Background(), storage.AllUsers, storage.RoleReader)
// 	if err != nil {
// 		return err
// 	}
// 	d.dumpDefaultACL()
// 	return nil
// }

// func (d *demo) deleteDefaultACLRule() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")
// 	w := &bytes.Buffer{}

// 	io.WriteString(w, "\nDelete Default object ACL Rule:\n")
// 	err = bucket.DefaultObjectACL().Delete(context.Background(), storage.AllUsers)
// 	if err != nil {
// 		return err
// 	}
// 	d.dumpDefaultACL()
// 	return nil
// }

// func (d *demo) dumpBucketACL() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")
// 	w := &bytes.Buffer{}
// 	acl, err := bucket.ACL().List(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	for _, v := range acl {
// 		fmt.Fprintf(w, "Scope: %q, Permission: %q\n", v.Entity, v.Role)
// 	}
// 	return nil
// }

// func (d *demo) bucketACL() {
// 	w := &bytes.Buffer{}
// 	io.WriteString(w, "\nBucket ACL:\n")
// 	d.dumpBucketACL()
// }

// func (d *demo) putBucketACLRule() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	w := &bytes.Buffer{}

// 	io.WriteString(w, "\nPut Bucket ACL Rule:\n")
// 	err = bucket.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader)
// 	if err != nil {
// 		return err
// 	}
// 	d.dumpBucketACL()
// 	return nil
// }

// func (d *demo) deleteBucketACLRule() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")
// 	w := &bytes.Buffer{}

// 	io.WriteString(w, "\nDelete Bucket ACL Rule:\n")
// 	err = bucket.ACL().Delete(context.Background(), storage.AllUsers)
// 	if err != nil {
// 		return err
// 	}
// 	d.dumpBucketACL()
// 	return nil
// }

// func (d *demo) dumpACL(fileName string) error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	acl, err := bucket.Object(fileName).ACL().List(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	for _, v := range acl {
// 		log.Println("Scope: %q, Permission: %q\n", v.Entity, v.Role)
// 	}
// 	return nil
// }

// func (d *demo) acl(fileName string) {
// 	log.Println("\nACL for file %v:\n", fileName)
// 	d.dumpACL(fileName)
// }

// func (d *demo) putACLRule(fileName string) error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	log.Println("\nPut ACL rule for file %v:\n", fileName)
// 	err = bucket.Object(fileName).ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader)
// 	if err != nil {
// 		return err
// 	}
// 	d.dumpACL(fileName)
// 	return nil
// }

// func (d *demo) deleteACLRule(fileName string) error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	log.Printf("\nDelete ACL rule for file %v:\n", fileName)
// 	err = bucket.Object(fileName).ACL().Delete(context.Background(), storage.AllUsers)
// 	if err != nil {
// 		return err
// 	}
// 	d.dumpACL(fileName)
// 	return nil
// }

// func (d *demo) deleteFiles() error {
// 	client, err := storage.NewClient(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer client.Close()

// 	bucket := client.Bucket("")

// 	log.Println("\nDeleting files...\n")
// 	for _, v := range d.cleanUp {
// 		log.Println("Deleting file %v\n", v)
// 		if err := bucket.Object(v).Delete(context.Background()); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
