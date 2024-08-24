// Copyright 2016 - 2024 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.18 or later.

package excelize

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/richardlehane/mscfb"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	// Test decrypt spreadsheet with incorrect password
	_, err := OpenFile(filepath.Join("test", "encryptSHA1.xlsx"), Options{Password: "passwd"})
	assert.EqualError(t, err, ErrWorkbookPassword.Error())
	// Test decrypt spreadsheet with password
	f, err := OpenFile(filepath.Join("test", "encryptSHA1.xlsx"), Options{Password: "password"})
	assert.NoError(t, err)
	cell, err := f.GetCellValue("Sheet1", "A1")
	assert.NoError(t, err)
	assert.Equal(t, "SECRET", cell)
	assert.NoError(t, f.Close())
	// Test decrypt spreadsheet with unsupported encrypt mechanism
	raw, err := os.ReadFile(filepath.Join("test", "encryptAES.xlsx"))
	assert.NoError(t, err)
	raw[2050] = 3
	_, err = Decrypt(raw, &Options{Password: "password"})
	assert.Equal(t, ErrUnsupportedEncryptMechanism, err)

	// Test encrypt spreadsheet with invalid password
	assert.EqualError(t, f.SaveAs(filepath.Join("test", "Encryption.xlsx"), Options{Password: strings.Repeat("*", MaxFieldLength+1)}), ErrPasswordLengthInvalid.Error())
	// Test encrypt spreadsheet with new password
	assert.NoError(t, f.SaveAs(filepath.Join("test", "Encryption.xlsx"), Options{Password: "passwd"}))
	assert.NoError(t, f.Close())
	f, err = OpenFile(filepath.Join("test", "Encryption.xlsx"), Options{Password: "passwd"})
	assert.NoError(t, err)
	cell, err = f.GetCellValue("Sheet1", "A1")
	assert.NoError(t, err)
	assert.Equal(t, "SECRET", cell)
	// Test remove password by save workbook with options
	assert.NoError(t, f.Save(Options{Password: ""}))
	assert.NoError(t, f.Close())

	doc, err := mscfb.New(bytes.NewReader(raw))
	assert.NoError(t, err)
	encryptionInfoBuf, encryptedPackageBuf := extractPart(doc)
	binary.LittleEndian.PutUint64(encryptionInfoBuf[20:32], uint64(0))
	_, err = standardDecrypt(encryptionInfoBuf, encryptedPackageBuf, &Options{Password: "password"})
	assert.NoError(t, err)
	_, err = decrypt(nil, nil, nil)
	assert.EqualError(t, err, "crypto/aes: invalid key size 0")
	_, err = agileDecrypt(encryptionInfoBuf, MacintoshCyrillicCharset, &Options{Password: "password"})
	assert.EqualError(t, err, "XML syntax error on line 1: invalid character entity &0 (no semicolon)")
	_, err = convertPasswdToKey("password", nil, Encryption{
		KeyEncryptors: KeyEncryptors{KeyEncryptor: []KeyEncryptor{
			{EncryptedKey: EncryptedKey{KeyData: KeyData{SaltValue: "=="}}},
		}},
	})
	assert.EqualError(t, err, "illegal base64 data at input byte 0")
	_, err = createIV([]byte{0}, Encryption{KeyData: KeyData{SaltValue: "=="}})
	assert.EqualError(t, err, "illegal base64 data at input byte 0")
}

func TestEncryptionMechanism(t *testing.T) {
	mechanism, err := encryptionMechanism([]byte{3, 0, 3, 0})
	assert.Equal(t, mechanism, "extensible")
	assert.EqualError(t, err, ErrUnsupportedEncryptMechanism.Error())
	_, err = encryptionMechanism([]byte{})
	assert.EqualError(t, err, ErrUnknownEncryptMechanism.Error())
}

func TestHashing(t *testing.T) {
	assert.Equal(t, hashing("unsupportedHashAlgorithm", []byte{}), []byte(nil))
}

func TestGenISOPasswdHash(t *testing.T) {
	for hashAlgorithm, expected := range map[string][]string{
		"MD4":     {"2lZQZUubVHLm/t6KsuHX4w==", "TTHjJdU70B/6Zq83XGhHVA=="},
		"MD5":     {"HWbqyd4dKKCjk1fEhk2kuQ==", "8ADyorkumWCayIukRhlVKQ=="},
		"SHA-1":   {"XErQIV3Ol+nhXkyCxrLTEQm+mSc=", "I3nDtyf59ASaNX1l6KpFnA=="},
		"SHA-256": {"7oqMFyfED+mPrzRIBQ+KpKT4SClMHEPOZldliP15xAA=", "ru1R/w3P3Jna2Qo+EE8QiA=="},
		"SHA-384": {"nMODLlxsC8vr0btcq0kp/jksg5FaI3az5Sjo1yZk+/x4bFzsuIvpDKUhJGAk/fzo", "Zjq9/jHlgOY6MzFDSlVNZg=="},
		"SHA-512": {"YZ6jrGOFQgVKK3rDK/0SHGGgxEmFJglQIIRamZc2PkxVtUBp54fQn96+jVXEOqo6dtCSanqksXGcm/h3KaiR4Q==", "p5s/bybHBPtusI7EydTIrg=="},
	} {
		hashValue, saltValue, err := genISOPasswdHash("password", hashAlgorithm, expected[1], int(sheetProtectionSpinCount))
		assert.NoError(t, err)
		assert.Equal(t, expected[0], hashValue)
		assert.Equal(t, expected[1], saltValue)
	}
}