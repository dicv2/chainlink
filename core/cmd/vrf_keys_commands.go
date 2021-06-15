package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/web/presenters"
	"go.uber.org/multierr"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/urfave/cli"
)

type VRFKeyPresenter struct {
	JAID // Include this to overwrite the presenter JAID so it can correctly render the ID in JSON
	presenters.VRFKeyResource
}

func (p VRFKeyPresenter) FriendlyDeletedAt() string {
	if p.DeletedAt != nil {
		return p.DeletedAt.String()
	}
	return ""
}

// RenderTable implements TableRenderer
func (p *VRFKeyPresenter) RenderTable(rt RendererTable) error {
	headers := []string{"Compressed", "Uncompressed", "Hash", "Created", "Updated", "Deleted"}
	rows := [][]string{p.ToRow()}
	renderList(headers, rows, rt.Writer)
	return nil
}

func (p *VRFKeyPresenter) ToRow() []string {
	return []string{
		p.Compressed,
		p.Uncompressed,
		p.Hash,
		p.CreatedAt.String(),
		p.UpdatedAt.String(),
		p.FriendlyDeletedAt(),
	}
}

type VRFKeyPresenters []VRFKeyPresenter

// RenderTable implements TableRenderer
func (ps VRFKeyPresenters) RenderTable(rt RendererTable) error {
	headers := []string{"Compressed", "Uncompressed", "Hash", "Created", "Updated", "Deleted"}
	rows := [][]string{}

	for _, p := range ps {
		rows = append(rows, p.ToRow())
	}

	renderList(headers, rows, rt.Writer)

	return nil
}

// CreateVRFKey creates a key in the VRF keystore, protected by the password in
// the password file
func (cli *Client) CreateVRFKey(c *cli.Context) error {
	resp, err := cli.HTTP.Post("/v2/keys/vrf", nil)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	var presenter VRFKeyPresenter
	return cli.renderAPIResponse(resp, &presenter)
}

// CreateAndExportWeakVRFKey creates a key in the VRF keystore, protected by the
// password in the password file, but with weak key-derivation-function
// parameters, which makes it cheaper for testing, but also more vulnerable to
// bruteforcing of the encrypted key material. For testing purposes only!
//
// The key is only stored at the specified file location, not stored in the DB.
func (cli *Client) CreateAndExportWeakVRFKey(c *cli.Context) error {
	password, err := getPassword(c)
	if err != nil {
		return err
	}
	app, err := cli.AppFactory.NewApplication(cli.Config)
	if err != nil {
		return cli.errorOut(errors.Wrap(err, "creating application"))
	}
	vrfKeyStore := app.GetKeyStore().VRF()
	key, err := vrfKeyStore.CreateAndUnlockWeakInMemoryEncryptedKeyXXXTestingOnly(
		string(password))
	if err != nil {
		return errors.Wrapf(err, "while creating testing key")
	}
	if !c.IsSet("file") || !noFileToOverwrite(c.String("file")) {
		errmsg := "must specify path to key file which does not already exist"
		fmt.Println(errmsg)
		return fmt.Errorf(errmsg)
	}
	fmt.Println("Don't use this key for anything sensitive!")
	return key.WriteToDisk(c.String("file"))
}

// ImportVRFKey reads a file into an EncryptedVRFKey in the db
func (cli *Client) ImportVRFKey(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("Must pass the filepath of the key to be imported"))
	}

	oldPasswordFile := c.String("oldpassword")
	if len(oldPasswordFile) == 0 {
		return cli.errorOut(errors.New("Must specify --oldpassword/-p flag"))
	}
	oldPassword, err := ioutil.ReadFile(oldPasswordFile)
	if err != nil {
		return cli.errorOut(errors.Wrap(err, "Could not read password file"))
	}

	filepath := c.Args().Get(0)
	keyJSON, err := ioutil.ReadFile(filepath)
	if err != nil {
		return cli.errorOut(err)
	}

	normalizedPassword := normalizePassword(string(oldPassword))
	resp, err := cli.HTTP.Post("/v2/keys/vrf/import?oldpassword="+normalizedPassword, bytes.NewReader(keyJSON))
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	var presenter VRFKeyPresenter
	return cli.renderAPIResponse(resp, &presenter, "Imported VRF key")
}

// ExportVRFKey saves encrypted copy of VRF key with given public key to
// requested file path.
func (cli *Client) ExportVRFKey(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("Must pass the ID of the key to export"))
	}

	newPasswordFile := c.String("newpassword")
	if len(newPasswordFile) == 0 {
		return cli.errorOut(errors.New("Must specify --newpassword/-p flag"))
	}
	newPassword, err := ioutil.ReadFile(newPasswordFile)
	if err != nil {
		return cli.errorOut(errors.Wrap(err, "Could not read password file"))
	}

	filepath := c.String("output")
	if len(filepath) == 0 {
		return cli.errorOut(errors.New("Must specify --output/-o flag"))
	}

	pk, err := getPublicKey(c)
	if err != nil {
		return cli.errorOut(err)
	}

	normalizedPassword := normalizePassword(string(newPassword))
	resp, err := cli.HTTP.Post("/v2/keys/ocr/export/"+pk.String()+"?newpassword="+normalizedPassword, nil)
	if err != nil {
		return cli.errorOut(errors.Wrap(err, "Could not make HTTP request"))
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return cli.errorOut(errors.New("Error exporting"))
	}

	keyJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cli.errorOut(errors.Wrap(err, "Could not read response body"))
	}

	err = utils.WriteFileWithMaxPerms(filepath, keyJSON, 0600)
	if err != nil {
		return cli.errorOut(errors.Wrapf(err, "Could not write %v", filepath))
	}

	_, err = os.Stderr.WriteString(fmt.Sprintf("Exported VRF key %s to %s", pk.String(), filepath))
	if err != nil {
		return cli.errorOut(err)
	}

	return nil
	/*
		encryptedKey, err := getKeys(cli, c)
		if err != nil {
			return err
		}
		if c.String("file") == "" {
			return fmt.Errorf("must specify file to export to") // Or could default to stdout?
		}
		keypath := c.String("file")
		_, err = os.Stat(keypath)
		if err == nil {
			return fmt.Errorf(
				"refusing to overwrite existing file %s. Please move it or change the save path",
				keypath)
		}
		if !os.IsNotExist(err) {
			return errors.Wrapf(err, "while checking whether file %s exists", keypath)
		}
		if err := encryptedKey.WriteToDisk(keypath); err != nil {
			return errors.Wrapf(err, "could not save %#+v to %s", encryptedKey, keypath)
		}

		return nil
	*/
}

// DeleteVRFKey soft-deletes the VRF key with given public key from the db
//
// Since this runs in an independent process from any chainlink node, it cannot
// cause running nodes to forget the key, if they already have it unlocked.
func (cli *Client) DeleteVRFKey(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("Must pass the key ID to be deleted"))
	}
	id, err := getPublicKey(c)
	if err != nil {
		return cli.errorOut(err)
	}

	if !confirmAction(c) {
		return nil
	}

	var queryStr string
	if c.Bool("hard") {
		queryStr = "?hard=true"
	}

	resp, err := cli.HTTP.Delete(fmt.Sprintf("/v2/keys/vrf/%s%s", id, queryStr))
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	var presenter VRFKeyPresenter
	return cli.renderAPIResponse(resp, &presenter, "OCR key bundle deleted")
}

func getPublicKey(c *cli.Context) (secp256k1.PublicKey, error) {
	if c.String("publicKey") == "" {
		return secp256k1.PublicKey{}, fmt.Errorf("must specify public key")
	}
	publicKey, err := secp256k1.NewPublicKeyFromHex(c.String("publicKey"))
	if err != nil {
		return secp256k1.PublicKey{}, errors.Wrap(err, "failed to parse public key")
	}
	return publicKey, nil
}

// ListKeys Lists the keys in the db
func (cli *Client) ListVRFKeys(c *cli.Context) error {
	resp, err := cli.HTTP.Get("/v2/keys/vrf", nil)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	var presenters VRFKeyPresenters
	return cli.renderAPIResponse(resp, &presenters, "🔑 VRF Keys")
}

func noFileToOverwrite(path string) bool {
	return os.IsNotExist(utils.JustError(os.Stat(path)))
}
