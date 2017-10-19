require 'serverspec'
require 'docker'

set :backend, :docker

class Specinfra::Command::Busybox < Specinfra::Command::Linux::Base
end
