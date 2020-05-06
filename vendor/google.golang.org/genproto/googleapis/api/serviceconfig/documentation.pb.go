// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/documentation.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// `Documentation` provides the information for describing a service.
//
// Example:
// <pre><code>documentation:
//   summary: >
//     The Google Calendar API gives access
//     to most calendar features.
//   pages:
//   - name: Overview
//     content: &#40;== include google/foo/overview.md ==&#41;
//   - name: Tutorial
//     content: &#40;== include google/foo/tutorial.md ==&#41;
//     subpages;
//     - name: Java
//       content: &#40;== include google/foo/tutorial_java.md ==&#41;
//   rules:
//   - selector: google.calendar.Calendar.Get
//     description: >
//       ...
//   - selector: google.calendar.Calendar.Put
//     description: >
//       ...
// </code></pre>
// Documentation is provided in markdown syntax. In addition to
// standard markdown features, definition lists, tables and fenced
// code blocks are supported. Section headers can be provided and are
// interpreted relative to the section nesting of the context where
// a documentation fragment is embedded.
//
// Documentation from the IDL is merged with documentation defined
// via the config at normalization time, where documentation provided
// by config rules overrides IDL provided.
//
// A number of constructs specific to the API platform are supported
// in documentation text.
//
// In order to reference a proto element, the following
// notation can be used:
// <pre><code>&#91;fully.qualified.proto.name]&#91;]</code></pre>
// To override the display text used for the link, this can be used:
// <pre><code>&#91;display text]&#91;fully.qualified.proto.name]</code></pre>
// Text can be excluded from doc using the following notation:
// <pre><code>&#40;-- internal comment --&#41;</code></pre>
//
// A few directives are available in documentation. Note that
// directives must appear on a single line to be properly
// identified. The `include` directive includes a markdown file from
// an external source:
// <pre><code>&#40;== include path/to/file ==&#41;</code></pre>
// The `resource_for` directive marks a message to be the resource of
// a collection in REST view. If it is not specified, tools attempt
// to infer the resource from the operations in a collection:
// <pre><code>&#40;== resource_for v1.shelves.books ==&#41;</code></pre>
// The directive `suppress_warning` does not directly affect documentation
// and is documented together with service config validation.
type Documentation struct {
	// A short summary of what the service does. Can only be provided by
	// plain text.
	Summary string `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	// The top level pages for the documentation set.
	Pages []*Page `protobuf:"bytes,5,rep,name=pages,proto3" json:"pages,omitempty"`
	// A list of documentation rules that apply to individual API elements.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*DocumentationRule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	// The URL to the root of documentation.
	DocumentationRootUrl string `protobuf:"bytes,4,opt,name=documentation_root_url,json=documentationRootUrl,proto3" json:"documentation_root_url,omitempty"`
	// Declares a single overview page. For example:
	// <pre><code>documentation:
	//   summary: ...
	//   overview: &#40;== include overview.md ==&#41;
	// </code></pre>
	// This is a shortcut for the following declaration (using pages style):
	// <pre><code>documentation:
	//   summary: ...
	//   pages:
	//   - name: Overview
	//     content: &#40;== include overview.md ==&#41;
	// </code></pre>
	// Note: you cannot specify both `overview` field and `pages` field.
	Overview             string   `protobuf:"bytes,2,opt,name=overview,proto3" json:"overview,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Documentation) Reset()         { *m = Documentation{} }
func (m *Documentation) String() string { return proto.CompactTextString(m) }
func (*Documentation) ProtoMessage()    {}
func (*Documentation) Descriptor() ([]byte, []int) {
	return fileDescriptor_documentation_3d6f0d592b1e580f, []int{0}
}
func (m *Documentation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Documentation.Unmarshal(m, b)
}
func (m *Documentation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Documentation.Marshal(b, m, deterministic)
}
func (dst *Documentation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Documentation.Merge(dst, src)
}
func (m *Documentation) XXX_Size() int {
	return xxx_messageInfo_Documentation.Size(m)
}
func (m *Documentation) XXX_DiscardUnknown() {
	xxx_messageInfo_Documentation.DiscardUnknown(m)
}

var xxx_messageInfo_Documentation proto.InternalMessageInfo

func (m *Documentation) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Documentation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *Documentation) GetRules() []*DocumentationRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Documentation) GetDocumentationRootUrl() string {
	if m != nil {
		return m.DocumentationRootUrl
	}
	return ""
}

func (m *Documentation) GetOverview() string {
	if m != nil {
		return m.Overview
	}
	return ""
}

// A documentation rule provides information about individual API elements.
type DocumentationRule struct {
	// The selector is a comma-separated list of patterns. Each pattern is a
	// qualified name of the element which may end in "*", indicating a wildcard.
	// Wildcards are only allowed at the end and for a whole component of the
	// qualified name, i.e. "foo.*" is ok, but not "foo.b*" or "foo.*.bar". To
	// specify a default for all applicable elements, the whole pattern "*"
	// is used.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Description of the selected API(s).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Deprecation description of the selected element(s). It can be provided if
	// an element is marked as `deprecated`.
	DeprecationDescription string   `protobuf:"bytes,3,opt,name=deprecation_description,json=deprecationDescription,proto3" json:"deprecation_description,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *DocumentationRule) Reset()         { *m = DocumentationRule{} }
func (m *DocumentationRule) String() string { return proto.CompactTextString(m) }
func (*DocumentationRule) ProtoMessage()    {}
func (*DocumentationRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_documentation_3d6f0d592b1e580f, []int{1}
}
func (m *DocumentationRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentationRule.Unmarshal(m, b)
}
func (m *DocumentationRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentationRule.Marshal(b, m, deterministic)
}
func (dst *DocumentationRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentationRule.Merge(dst, src)
}
func (m *DocumentationRule) XXX_Size() int {
	return xxx_messageInfo_DocumentationRule.Size(m)
}
func (m *DocumentationRule) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentationRule.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentationRule proto.InternalMessageInfo

func (m *DocumentationRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *DocumentationRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DocumentationRule) GetDeprecationDescription() string {
	if m != nil {
		return m.DeprecationDescription
	}
	return ""
}

// Represents a documentation page. A page can contain subpages to represent
// nested documentation set structure.
type Page struct {
	// The name of the page. It will be used as an identity of the page to
	// generate URI of the page, text of the link to this page in navigation,
	// etc. The full page name (start from the root page name to this page
	// concatenated with `.`) can be used as reference to the page in your
	// documentation. For example:
	// <pre><code>pages:
	// - name: Tutorial
	//   content: &#40;== include tutorial.md ==&#41;
	//   subpages:
	//   - name: Java
	//     content: &#40;== include tutorial_java.md ==&#41;
	// </code></pre>
	// You can reference `Java` page using Markdown reference link syntax:
	// `[Java][Tutorial.Java]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Markdown content of the page. You can use <code>&#40;== include {path}
	// ==&#41;</code> to include content from a Markdown file.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Subpages of this page. The order of subpages specified here will be
	// honored in the generated docset.
	Subpages             []*Page  `protobuf:"bytes,3,rep,name=subpages,proto3" json:"subpages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_documentation_3d6f0d592b1e580f, []int{2}
}
func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (dst *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(dst, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Page) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Page) GetSubpages() []*Page {
	if m != nil {
		return m.Subpages
	}
	return nil
}

func init() {
	proto.RegisterType((*Documentation)(nil), "google.api.Documentation")
	proto.RegisterType((*DocumentationRule)(nil), "google.api.DocumentationRule")
	proto.RegisterType((*Page)(nil), "google.api.Page")
}

func init() {
	proto.RegisterFile("google/api/documentation.proto", fileDescriptor_documentation_3d6f0d592b1e580f)
}

var fileDescriptor_documentation_3d6f0d592b1e580f = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6a, 0xe3, 0x30,
	0x14, 0x45, 0x71, 0xec, 0xcc, 0x64, 0x5e, 0x98, 0x61, 0x46, 0x0c, 0x19, 0x33, 0xd0, 0x12, 0xb2,
	0x28, 0x59, 0x14, 0x1b, 0x9a, 0x42, 0x17, 0x5d, 0x35, 0xa4, 0x94, 0xee, 0x8c, 0xa1, 0x9b, 0x6e,
	0x82, 0xa2, 0xbc, 0x0a, 0x83, 0xad, 0x67, 0x24, 0x39, 0xa5, 0xbf, 0xd0, 0xcf, 0xe8, 0x57, 0xf5,
	0x73, 0x8a, 0x65, 0x27, 0xb1, 0x29, 0xdd, 0xf9, 0xfa, 0x1e, 0xe9, 0x3e, 0x5d, 0x09, 0x4e, 0x25,
	0x91, 0xcc, 0x31, 0xe6, 0x65, 0x16, 0x6f, 0x49, 0x54, 0x05, 0x2a, 0xcb, 0x6d, 0x46, 0x2a, 0x2a,
	0x35, 0x59, 0x62, 0xd0, 0xf8, 0x11, 0x2f, 0xb3, 0xd9, 0xbb, 0x07, 0x3f, 0x57, 0x5d, 0x86, 0x85,
	0xf0, 0xdd, 0x54, 0x45, 0xc1, 0xf5, 0x4b, 0xe8, 0x4d, 0xbd, 0xf9, 0x8f, 0x74, 0x2f, 0xd9, 0x19,
	0x0c, 0x4b, 0x2e, 0xd1, 0x84, 0xc3, 0xa9, 0x3f, 0x1f, 0x5f, 0xfc, 0x8e, 0x8e, 0xfb, 0x44, 0x09,
	0x97, 0x98, 0x36, 0x36, 0x5b, 0xc0, 0x50, 0x57, 0x39, 0x9a, 0xd0, 0x77, 0xdc, 0x49, 0x97, 0xeb,
	0x65, 0xa5, 0x55, 0x8e, 0x69, 0xc3, 0xb2, 0x4b, 0x98, 0xf4, 0x66, 0x5d, 0x6b, 0x22, 0xbb, 0xae,
	0x74, 0x1e, 0x06, 0x6e, 0x8a, 0xbf, 0x3d, 0x37, 0x25, 0xb2, 0x0f, 0x3a, 0x67, 0xff, 0x61, 0x44,
	0x3b, 0xd4, 0xbb, 0x0c, 0x9f, 0xc3, 0x81, 0xe3, 0x0e, 0x7a, 0xf6, 0xea, 0xc1, 0x9f, 0x4f, 0x71,
	0xf5, 0x0a, 0x83, 0x39, 0x0a, 0x4b, 0xba, 0x3d, 0xdf, 0x41, 0xb3, 0x29, 0x8c, 0xb7, 0x68, 0x84,
	0xce, 0xca, 0x1a, 0x6f, 0x37, 0xec, 0xfe, 0x62, 0x57, 0xf0, 0x6f, 0x8b, 0xa5, 0x46, 0xd1, 0xcc,
	0xd8, 0xa5, 0x7d, 0x47, 0x4f, 0x3a, 0xf6, 0xea, 0xe8, 0xce, 0x36, 0x10, 0xd4, 0x15, 0x31, 0x06,
	0x81, 0xe2, 0x05, 0xb6, 0xd1, 0xee, 0xbb, 0x6e, 0x5c, 0x90, 0xb2, 0xa8, 0x6c, 0x1b, 0xb9, 0x97,
	0xec, 0x1c, 0x46, 0xa6, 0xda, 0x34, 0xa5, 0xfb, 0x5f, 0x94, 0x7e, 0x20, 0x96, 0x16, 0x7e, 0x09,
	0x2a, 0x3a, 0xc0, 0x92, 0xf5, 0xce, 0x9f, 0xd4, 0xb7, 0x9f, 0x78, 0x8f, 0xb7, 0x2d, 0x21, 0x29,
	0xe7, 0x4a, 0x46, 0xa4, 0x65, 0x2c, 0x51, 0xb9, 0xb7, 0x11, 0x37, 0x16, 0x2f, 0x33, 0xe3, 0x9e,
	0x8f, 0xa9, 0xbb, 0x14, 0x28, 0x48, 0x3d, 0x65, 0xf2, 0xba, 0xa7, 0xde, 0x06, 0xc1, 0xdd, 0x4d,
	0x72, 0xbf, 0xf9, 0xe6, 0x16, 0x2e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x04, 0x32, 0xbf,
	0x76, 0x02, 0x00, 0x00,
}
