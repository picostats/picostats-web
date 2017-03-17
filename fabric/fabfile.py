#!/usr/bin/env python
from fabric.api import env, local, run, sudo
from fabric.operations import put

env.hosts = ['picostats@46.101.254.59']


def deploy():
    local('mkdir picostatsnew-web')
    local('cp ../picostats-web picostatsnew-web/')
    local('cp -R ../public picostatsnew-web/')
    local('cp -R ../templates picostatsnew-web/')
    local('cp production.json picostatsnew-web/config.json')
    local('tar -czf picostatsnew-web.tar.gz picostatsnew-web/')
    local('rm -rf picostatsnew-web')
    put('picostatsnew-web.tar.gz')
    run('tar -xzf picostatsnew-web.tar.gz')
    run('rm -rf picostats-web')
    run('mv picostatsnew-web picostats-web')
    run('rm picostatsnew-web.tar.gz')
    sudo('supervisorctl reload')
