package jsonpatch

import (
	"encoding/json"
	//"github.com/bitly/go-simplejson"
)

/*
Patch is a list of Patch Operations.

>>> patch = JsonPatch([
...     {'op': 'add', 'path': '/foo', 'value': 'bar'},
...     {'op': 'add', 'path': '/baz', 'value': [1, 2, 3]},
...     {'op': 'remove', 'path': '/baz/1'},
...     {'op': 'test', 'path': '/baz', 'value': [1, 3]},
...     {'op': 'replace', 'path': '/baz/0', 'value': 42},
...     {'op': 'remove', 'path': '/baz/1'},
... ])
>>> doc = {}
>>> result = patch.apply(doc)
>>> expected = {'foo': 'bar', 'baz': [42]}
>>> result == expected
True

JsonPatch object is iterable, so you could easily access to each patch
statement in loop:

>>> lpatch = list(patch)
>>> expected = {'op': 'add', 'path': '/foo', 'value': 'bar'}
>>> lpatch[0] == expected
True
>>> lpatch == patch.patch
True

Also JsonPatch could be converted directly to :class:`bool` if it contains
any operation statements:

>>> bool(patch)
True
>>> bool(JsonPatch([]))
False

This behavior is very handy with :func:`make_patch` to write more readable
code:

>>> old = {'foo': 'bar', 'numbers': [1, 3, 4, 8]}
>>> new = {'baz': 'qux', 'numbers': [1, 4, 7]}
>>> patch = make_patch(old, new)
>>> if patch:
...     # document have changed, do something useful
...     patch.apply(old)    #doctest: +ELLIPSIS
{...}
*/
type Patch struct {
	Operations []PatchOperation
}

func (p *Patch) Apply(doc interface{}) error { return nil }

func FromString(str string) (PatchOperation, error) {
	patch := PatchOperation{}
	err := json.Unmarshal([]byte(str), &patch)
	return patch, err
}

func FromDiff(src interface{}, dst interface{}) (Patch, error) {
	return Patch{}, nil
}

/*
MakePatch generates patch by comparing of two document objects. Actually is
a proxy to :meth:`JsonPatch.from_diff` method.

:param src: Data source document object.
:type src: dict

:param dst: Data source document object.
:type dst: dict

>>> src = {'foo': 'bar', 'numbers': [1, 3, 4, 8]}
>>> dst = {'baz': 'qux', 'numbers': [1, 4, 7]}
>>> patch = make_patch(src, dst)
>>> new = patch.apply(src)
>>> new == dst
True
*/
func MakePatch(src interface{}, dst interface{}) (Patch, error) {
	return FromDiff(src, dst)
}
