require 'serverspec'
require 'docker'

set :backend, :docker
set :os, :family => 'linux'
