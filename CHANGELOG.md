# Changelog

## [0.2.0](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.1.2...v0.2.0) (2023-01-16)


### ⚠ BREAKING CHANGES

* **vpc:** remove floatingIPAssociate resource
* **vpc:** remove network resource
* **vpc:** remove networkingSubnet resource
* **vpc:** remove router resource
* **vpc:** remove routerInterface resource
* **vpc:** remove networking resources ([#91](https://github.com/FrangipaneTeam/provider-flexibleengine/issues/91))
* **main:** add tenantName to ProviderConfig.spec
* **vpc/vip:** remove deprecated subnet_id argument

### Features

* **dcs:** Add dcs examples ([281ff49](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/281ff4931f73f859b1dc3b6ee6358201e59bdd0e))
* **dms:** Add dms kafka instance example ([c3d3bb2](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **dms:** Add dms kafka topic and user examples ([6a84569](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/6a845690deff1489683ce34e640bcbdba1ce3d7e))
* **dms:** Add kafka instance example ([c3d3bb2](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **ecs:** add floatingIpAssociate example ([2c3f301](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/2c3f3013833e143d0fdfb1ef9552b0366e3be109))
* **ecs:** add interfaceAttach example ([0040d48](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/0040d4831fc9f2559ff7e95faf96c6d1b2807e99))
* **evs/computevolumeattach:** add example EVS ComputeVolumeAttach ([c2a80a9](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/c2a80a9a741fca3c50fd1c497805efa8e6c9bf3b))
* **main:** add domainId and tenantId ([c30de00](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/c30de0034ee7654d85366dbf9f899120d02e3fa7))
* **main:** add tenantName to ProviderConfig.spec ([beb5c2a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/beb5c2ae1f20337cc9d64a7e2121efe4b1081457))
* **vpc:** add examples ([#74](https://github.com/FrangipaneTeam/provider-flexibleengine/issues/74)) ([cbe31d3](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cbe31d3ed9aece7ed372be39db3e30aae09fe5b5))


### Bug Fixes

* **dcs:** add kindmap resource ([aa7e4f5](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/aa7e4f55b330f92ad89a054c275b22c02728d069))
* **dms:** fix instance_id ([3a17c73](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/3a17c733118f0dab2d0dc294793e86e1ff45765e))
* **dms:** Fix label example ([c3d3bb2](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **elb:** local type in global override ([#97](https://github.com/FrangipaneTeam/provider-flexibleengine/issues/97)) ([1e922e3](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/1e922e372be949ca7f34e858166d805ba7aa0bb0))
* **vpc/vip:** remove deprecated subnet_id argument ([44f769b](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/44f769b2dd29b768569fe1dfa82235abb4a2b629))
* **vpc:** fix subnet_id attribute ([8822745](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/88227451a6125ca8d63c6196c7fdddd5bff84da9))
* **vpc:** remove floatingIPAssociate resource ([cc1101a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove network resource ([cc1101a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove networking resources ([#91](https://github.com/FrangipaneTeam/provider-flexibleengine/issues/91)) ([cc1101a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove networkingSubnet resource ([cc1101a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove router resource ([cc1101a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove routerInterface resource ([cc1101a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))

## [0.1.2](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.1.1...v0.1.2) (2023-01-10)


### Features

* **dws:** add example cluster ([5b7c449](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/5b7c449f0accea9e87bf769a8221cdb8577e412c))
* **evs:** add availabilityZone in example ([053c702](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/053c702e1331000f3c7c4a42b464ca377ffb01f2))
* **rds:** add example ([86f463c](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** add example ([86f463c](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** Add example ([86f463c](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** add replicas example ([86f463c](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **sdrs:** add examples ([f63bbbd](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/f63bbbd16e08911941356846e120a3f4f7a6c0cd))


### Bug Fixes

* **dws:** example cluster nodetype and fix secret ([62c44d3](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/62c44d33ebf13bb9e0aacc99f53a387dcab91582))

## [0.1.1](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.1.0...v0.1.1) (2023-01-09)


### Bug Fixes

* **provider:** allow empty insecure parameter ([4a0c109](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/4a0c109e8ae1658f4b05627c4a685f8986589e5e))
* **provider:** allow empty insecure parameter ([d63d1a6](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/d63d1a6a455c9205beb0fca6adc42c171bf35a7e))

## [0.1.0](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.0.4...v0.1.0) (2023-01-09)


### ⚠ BREAKING CHANGES

* **cts:** Remove CTS tracker

### Features

* **cce:** Add CCE example ([bb84269](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/bb8426976652961dc5daded5e4f3ae988b1b9eec))
* **cts:** Remove CTS tracker ([462f740](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/462f7402161913a5ec16e4c8242dc9839694d9d2))
* **evs:** add examples ([a85d55e](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/a85d55e58facfb7b71f5ebbfe2fdf937a814a24f))
* **rts:** add examples ([8396380](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/8396380bedaf42c281b1b5400a7a02c71bf59acf))
* **vbs:** add examples ([7523159](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/75231598f3b773e4f985f6c31aacb3d30781f41c))
* **vbs:** add examples ([1b99a76](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/1b99a7632c4fd24d804d524c699e804d2fb16680))


### Bug Fixes

* **cce:** Fix CCE ([bb84269](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/bb8426976652961dc5daded5e4f3ae988b1b9eec))
* **cce:** Fix NameSpace ([8b19be6](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/8b19be6ae4241a183c4c3c2b10a5a80b5785f1ba))
* **evs:** Remove deprecated consistency_group_id ([3c3a27d](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/3c3a27d8e47693e2f179ffc504ae443624c450e1))
* **script:** fix bug check if label/ref is require ([8fd9b75](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/8fd9b759f2c7b76d81c3219a209dc638489d6e27))

## [0.0.4](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.0.3...v0.0.4) (2023-01-05)


### Features

* **dds:** add example instance ([24b8322](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/24b8322b672055c031946e819047d3df91368200))
* **dds:** fix example instance ([24b8322](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/24b8322b672055c031946e819047d3df91368200))
* **dis:** add example stream ([4ca4c80](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/4ca4c80164c23039ef00962743bc369d1b4a86c4))
* **dli:** add example database ([3380392](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/3380392e590ce9c1149e111c50198b5e18f610d2))
* **dli:** add example flink sql job ([c7b8d58](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/c7b8d584f231662e5d3c16a65c2931883dd9c410))
* **dli:** add example package ([1b82319](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/1b82319f8a23fe63a633296706e08f8f32205cec))
* **dli:** add example queue ([0f7ebcf](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/0f7ebcf0ddf8c349a731725e0fb544b04d6daebc))
* **dli:** add example spark job ([2223ac6](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/2223ac6e875a186ff828549dae15942faa81ec94))
* **dns:** add examples for zone/record/ptr ([#31](https://github.com/FrangipaneTeam/provider-flexibleengine/issues/31)) ([afce8ad](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/afce8add9a22b3e8691c183ce29490ee8fa42791))
* **eps:** add example project ([1e99b79](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/1e99b79e8a7aa6289c827cf7c402c333d87dad22))
* **kms:** add example key ([705a347](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/705a34752aec60504de4ac1d40bba93a55d2627a))
* **script:** add .extra files for kubectl generate command ([a64738f](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/a64738f34a965914b81390b0ecc537ae96ddd709))
* **script:** add support for SecretRef ([2f92dac](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/2f92daca909e6a589408e767286be9812df27e3b))
* **tms:** add example tags ([84e7be5](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/84e7be5db03cbbb2852927fcb8c59ba02bea401a))
* **vpcep:** add examples ([58a5194](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/58a5194266ef3aeb758148938d618084a8210a6d))
* **waf:** add Example Certificate ([8391107](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/83911078839e58dce1eb4be88ab385ccc42b76a9))
* **waf:** add Example Domain ([43e60eb](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/43e60eb1eba10bb0d537bb3df58cc8f5a5b45b09))
* **waf:** add Example Policy ([1d24f1b](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/1d24f1b1400384257bdabe6c7a9649b7651d3b09))
* **waf:** add Example RuleAlarmMaskings ([411d29a](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/411d29a6c8749afb841d864b54509bf4dd13caef))
* **waf:** add Example RuleBlackList ([419e137](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/419e1378aea9cf654047dc2d7e4188877fa03e44))
* **waf:** add Example RuleCcProtection ([84a46d4](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/84a46d42c70cbc1c834dd90b65e00c1f3a2a14ae))
* **waf:** add Example RuleDataMasking ([b01870d](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/b01870d3553be5cea0c5a37e0dba61cddcc34856))
* **waf:** add Example RulePreciseProtection ([97bc660](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/97bc66031ab962ee3deab76b4c6e90155ae21163))
* **waf:** add Example RuleWebTamperProtection ([363b5df](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/363b5df94078bc988db9a020cd18c595ec3a02be))
* **waf:** add Examples Dedicated ([caf2ff0](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/caf2ff0758aacb96b06832f5a45702b0b0a99582))


### Bug Fixes

* **dli:** minor fix for dli resources ([49a721f](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/49a721ffb8056839bece134fed3cbd66757b99bc))

## [0.0.3](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.0.2...v0.0.3) (2022-12-22)


### Features

* **lts:** add support and examples of Log Tank Service ([a7a2bbe](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/a7a2bbe97a6d7bde8da0333d1f8ad14fb6f12bc2))


### Bug Fixes

* mrs ([a3597ef](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/a3597efc3ba606cab2e666d6455c9af36f3e4310))
* **scripts:** bad `sys.argv` in Kubectl generate command ([d85ab0d](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/d85ab0d04a476354dd4111466b800ae0fe669c48))

## [0.0.2](https://github.com/FrangipaneTeam/provider-flexibleengine/compare/v0.0.1...v0.0.2) (2022-12-22)


### Features

* **lts:** add support and examples of Log Tank Service ([a7a2bbe](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/a7a2bbe97a6d7bde8da0333d1f8ad14fb6f12bc2))


### Bug Fixes

* **scripts:** bad `sys.argv` in Kubectl generate command ([d85ab0d](https://github.com/FrangipaneTeam/provider-flexibleengine/commit/d85ab0d04a476354dd4111466b800ae0fe669c48))
