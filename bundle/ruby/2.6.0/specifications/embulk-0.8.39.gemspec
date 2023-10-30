# -*- encoding: utf-8 -*-
# stub: embulk 0.8.39 ruby lib

Gem::Specification.new do |s|
  s.name = "embulk".freeze
  s.version = "0.8.39"

  s.required_rubygems_version = Gem::Requirement.new(">= 0".freeze) if s.respond_to? :required_rubygems_version=
  s.require_paths = ["lib".freeze]
  s.authors = ["Sadayuki Furuhashi".freeze]
  s.date = "2017-12-06"
  s.description = "Embulk is an open-source, plugin-based bulk data loader to scale and simplify data management across heterogeneous data stores. It can collect and ship any kinds of data in high throughput with transaction control.".freeze
  s.email = ["frsyuki@gmail.com".freeze]
  s.executables = ["embulk".freeze]
  s.files = ["bin/embulk".freeze]
  s.homepage = "https://github.com/embulk/embulk".freeze
  s.licenses = ["Apache 2.0".freeze]
  s.rubygems_version = "3.0.3.1".freeze
  s.summary = "Embulk, a plugin-based parallel bulk data loader".freeze

  s.installed_by_version = "3.0.3.1" if s.respond_to? :installed_by_version

  if s.respond_to? :specification_version then
    s.specification_version = 4

    if Gem::Version.new(Gem::VERSION) >= Gem::Version.new('1.2.0') then
      s.add_runtime_dependency(%q<jruby-jars>.freeze, ["= 9.1.13.0"])
      s.add_development_dependency(%q<rake>.freeze, [">= 0.10.0"])
      s.add_development_dependency(%q<test-unit>.freeze, ["~> 3.0.9"])
      s.add_development_dependency(%q<yard>.freeze, ["~> 0.8.7"])
      s.add_development_dependency(%q<kramdown>.freeze, ["~> 1.5.0"])
    else
      s.add_dependency(%q<jruby-jars>.freeze, ["= 9.1.13.0"])
      s.add_dependency(%q<rake>.freeze, [">= 0.10.0"])
      s.add_dependency(%q<test-unit>.freeze, ["~> 3.0.9"])
      s.add_dependency(%q<yard>.freeze, ["~> 0.8.7"])
      s.add_dependency(%q<kramdown>.freeze, ["~> 1.5.0"])
    end
  else
    s.add_dependency(%q<jruby-jars>.freeze, ["= 9.1.13.0"])
    s.add_dependency(%q<rake>.freeze, [">= 0.10.0"])
    s.add_dependency(%q<test-unit>.freeze, ["~> 3.0.9"])
    s.add_dependency(%q<yard>.freeze, ["~> 0.8.7"])
    s.add_dependency(%q<kramdown>.freeze, ["~> 1.5.0"])
  end
end
