package model

import (
	"fmt"
	"io"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
)

type MyURL struct {
	url.URL
}

func (u MyURL) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, u.URL.String()))
}

func (u *MyURL) UnmarshalGQL(v any) error {
	switch v := v.(type) {
	case string:
		result, err := url.Parse(v)
		if err != nil {
			return err
		}
		u = &MyURL{*result}
		return nil
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return err
		}
		u = &MyURL{*result}
		return nil
	default:
		return fmt.Errorf("%T is not a url.URL", v)
	}
}

func MarshalURI(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, u.String()))
	})
}

func UnmarshalURI(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		u, err := url.Parse(v)
		if err != nil {
			return url.URL{}, err
		}
		return *u, nil
	case []byte:
		u := &url.URL{}
		if err := u.UnmarshalBinary(v); err != nil {
			return url.URL{}, err
		}
		return *u, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a url.URL", v)
	}
}
