#
# spec file for whizz-ansible-modules package
#

%define whizz_collection /ansible_collections/whizz/embedded
%define architectures "amd64" "arm" "arm64"
%define modmap \
mkmodmap() {\
    local -n m=$1\
    m=(\
	[uname]="system/uname"\
	[ping]="system/ping"\
	[pkgpinch]="packaging/os/pkgpinch"\
	[zypper]="packaging/os/zypper"\
	[zypper_repository]="packaging/os/zypper_repository"\
    )\
}

Name:           whizz-ansible-modules
Version:        0.9
Release:        0
Summary:        Whizz Ansible-compatible collection of modules
License:        MIT
Group:          System/Tools
Url:            https://github.com/infra-whizz/wz-ansible-modules
Source:         %{name}-%{version}.tar.gz
Source1:        vendor.tar.gz

BuildRequires:  %{python_module py >= 1.4}
BuildRequires:  python3-base
BuildRequires:  python-rpm-macros
BuildRequires:  golang-packaging
BuildRequires:  golang(API) >= 1.13
Requires:       ansible

%description
Collection of Ansible-compatible modules, used by Whizz

%prep
%setup -q
%setup -q -T -D -a 1

%build
%modmap

declare -A modules
mkmodmap modules

# Compile binaries
for arch in %{architectures}
do
    for m_name in "${!modules[@]}"
    do
	CGO_ENABLED=0 GOOS=linux GOARCH=$arch go build -a -mod=vendor -tags netgo -ldflags '-w -extldflags "-static"' \
		   -o modules/${modules[$m_name]}/$m_name.$arch modules/${modules[$m_name]}/*.go
    done
done

%install
%modmap

mkdir -p %{buildroot}%{python3_sitelib}%{whizz_collection}

for d_name in "docs" "lib" "library" "plugins" "plugins/action" "roles"
do
    mkdir -p %{buildroot}%{python3_sitelib}%{whizz_collection}/$d_name
done

# Install generic collection structure
for f_name in "FILES.json" "MANIFEST.json" "README.md"
do
    install -m 0644 collection/whizz/embedded/$f_name %{buildroot}%{python3_sitelib}%{whizz_collection}
done

install -m 0644 collection/whizz/embedded/lib/* %{buildroot}%{python3_sitelib}%{whizz_collection}/lib

declare -A modules
mkmodmap modules

# Install pre-compiled binaries
for arch in %{architectures}
do
    for m_name in "${!modules[@]}"
    do
	if [ "$arch" = "amd64" ]; then
	    x_arch="x84_64"
	else
	    x_arch="$arch"
	fi
	install -m 0755 modules/${modules[$m_name]}/$m_name.$arch \
		%{buildroot}%{python3_sitelib}%{whizz_collection}/library/$m_name-linux-$x_arch
    done
done

# Install Python module wrappers
cd %{buildroot}%{python3_sitelib}%{whizz_collection}/plugins/action/
for m_name in "${!modules[@]}"
do
    ln -s ../../lib/basemodule.py $m_name.py
done

%files
%defattr(-,root,root)
%{python3_sitelib}%{whizz_collection}/*
%dir %{python3_sitelib}/ansible_collections/
%dir %{python3_sitelib}/ansible_collections/whizz/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/plugins/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/plugins/action/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/roles/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/lib/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/library/
%dir %{python3_sitelib}/ansible_collections/whizz/embedded/docs/

%changelog
