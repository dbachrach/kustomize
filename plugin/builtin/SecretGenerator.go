// Code generated by pluginator on SecretGenerator; DO NOT EDIT.
package builtin

import (
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/types"
	"sigs.k8s.io/yaml"
)

type SecretGeneratorPlugin struct {
	h                *resmap.PluginHelpers
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	types.GeneratorOptions
	types.SecretArgs
}

func (p *SecretGeneratorPlugin) Config(h *resmap.PluginHelpers, config []byte) (err error) {
	p.GeneratorOptions = types.GeneratorOptions{}
	p.SecretArgs = types.SecretArgs{}
	err = yaml.Unmarshal(config, p)
	if p.SecretArgs.Name == "" {
		p.SecretArgs.Name = p.Name
	}
	if p.SecretArgs.Namespace == "" {
		p.SecretArgs.Namespace = p.Namespace
	}
	p.h = h
	return
}

func (p *SecretGeneratorPlugin) Generate() (resmap.ResMap, error) {
	return p.h.ResmapFactory().FromSecretArgs(
		p.h.Loader(), &p.GeneratorOptions, p.SecretArgs)
}

func NewSecretGeneratorPlugin() resmap.GeneratorPlugin {
	return &SecretGeneratorPlugin{}
}
