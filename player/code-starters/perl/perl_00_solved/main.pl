#!/usr/bin/perl

use strict;
use warnings;
use aliased 'Javonet::Javonet' => 'Javonet';

Javonet->activate("p5XB-z7MN-Tp9a-d3NH-y4GA");

my $python_runtime = Javonet->in_memory()->python();
$python_runtime->load_library(".");

my $class_name = "robot-connector.Robot";
my $python_type = $python_runtime->get_type($class_name)->execute();

$python_type->invoke_static_method("solve")->execute();