// Copyright © 2019 Martin Tournoij – This file is part of GoatCounter and
// published under the terms of a slightly modified EUPL v1.2 license, which can
// be found in the LICENSE file or at https://license.goatcounter.com

package gctest

import (
	"fmt"
	"testing"

	"zgo.at/goatcounter"
	"zgo.at/zdb"
	"zgo.at/zlog"
	"zgo.at/zstd/zstring"
)

func TestDB(t *testing.T) {
	zlog.SetDebug("gctest")
	t.Run("", func(t *testing.T) {
		fmt.Println("Run 1")
		DB(t)
	})

	t.Run("", func(t *testing.T) {
		fmt.Println("\nRun 2")
		DB(t)
	})

	t.Run("", func(t *testing.T) {
		fmt.Println("\nRun 3")
		DB(t)
	})
}

func TestSite(t *testing.T) {
	ctx := DB(t)

	{
		var (
			s goatcounter.Site
			u goatcounter.User
		)
		Site(ctx, t, &s, &u)

		if s.ID == 0 {
			t.Error()
		}
		if u.ID == 0 {
			t.Error()
		}

		got := zdb.DumpString(ctx, `select site_id, cname from sites`) +
			zdb.DumpString(ctx, `select user_id, site_id, email from users`)
		want := `
			site_id  cname
			1        gctest.localhost
			2        NULL
			user_id  site_id  email
			1        1        test@gctest.localhost
			2        2        test@example.com`
		if d := zdb.Diff(got, want); d != "" {
			t.Error(d)
		}
	}

	{
		s := goatcounter.Site{Cname: zstring.NewPtr("XXX.XXX").P}
		u := goatcounter.User{Email: "XXX@XXX.XXX"}
		Site(ctx, t, &s, &u)

		if s.ID == 0 || *s.Cname != "XXX.XXX" {
			t.Error()
		}
		if u.ID == 0 || u.Email != "XXX@XXX.XXX" {
			t.Error()
		}

		got := zdb.DumpString(ctx, `select site_id, cname from sites`) +
			zdb.DumpString(ctx, `select user_id, site_id, email from users`)
		want := `
			site_id  cname
			1        gctest.localhost
			2        NULL
			3        XXX.XXX
			user_id  site_id  email
			1        1        test@gctest.localhost
			2        2        test@example.com
			3        3        XXX@XXX.XXX`
		if d := zdb.Diff(got, want); d != "" {
			t.Error(d)
		}
	}

	{
		Site(ctx, t, nil, nil)
		got := zdb.DumpString(ctx, `select site_id, cname from sites`) +
			zdb.DumpString(ctx, `select user_id, site_id, email from users`)
		want := `
			site_id  cname
			1        gctest.localhost
			2        NULL
			3        XXX.XXX
			4        NULL
			user_id  site_id  email
			1        1        test@gctest.localhost
			2        2        test@example.com
			3        3        XXX@XXX.XXX
			4        4        test@example.com`
		if d := zdb.Diff(got, want); d != "" {
			t.Error(d)
		}
	}
}
