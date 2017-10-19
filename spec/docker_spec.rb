require "spec_helper"

describe "Docker" do
  before(:all) do
    image = Docker::Image.build_from_dir('.')
    set :docker_image, image.id
  end

  describe file 'salamander' do
    it { should be_file }
  end

  describe port "8080" do
    it { should be_listening }
  end
end
