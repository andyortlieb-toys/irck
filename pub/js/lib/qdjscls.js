/* qdjscls.js
 * Quick & Dirty JS Class System
 *
 * @author Andrew Ortlieb <andyortlieb@gmail.com>
 * @website https://gist.github.com/andyortlieb/5051158 
 * @license WTFPL v2:
 *
 * What's the point?
 * Just a simple class system for your simple projects that could use
 * basic yet proper inheritance.
 * Load this file, or copy&paste the parts you want, or do whatever.
 *
 ********************************************************************
            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                   Version 2, December 2004

Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
  TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
 ********************************************************************
 *
 */
    /* TODO:
        - Come up with a class/superclass accessor shortcut system
    */

/* Library Definition.
 * You can also just copypasta the guts of this function if you'd rather.
 */
;(function(globals){
    "use strict";
    // The most basic level class 
    var BASE_CLASS = function(conf){
        var conf = conf || {};
        if (!conf.__NO_INIT__){
            this.conf = conf;
            for (var k in this.conf){
                this[k] = this.conf[k];
            }
            return this.init.apply(this,arguments);
        }
    };

    // init() may be overridden at the class, subclass or instance level.
    BASE_CLASS.prototype.init = function(){}; 

    BASE_CLASS.prototype.SuperAttr = function(type,attr){
        try{
            return type.prototype[attr];
        } catch (e){
            throw "`"+attr+"` is not available on superclass "+type+".";
        }
    };
    BASE_CLASS.prototype.SuperApply = function(type,attr){
        var fn = this.SuperAttr(type,attr);
        if (typeof(fn) !='function'){
            throw "`"+attr+"` is not a function on superclass "+type+".";
        }
        return fn.apply(this, Array.prototype.slice.call(arguments,2)[0]);
    };
    BASE_CLASS.prototype.SuperCall = function(type,attr){
        var fn = this.SuperAttr(type,attr);
        if (typeof(fn) !='function'){
            throw "`"+attr+"` is not a function on superclass"+type+".";
        }
        return fn.apply(this, Array.prototype.slice.call(arguments,2));
    };

    BASE_CLASS.prototype.ClassAttr = function(attr){
        try{
            return this._class.prototype[attr];
        } catch (e){
            throw "`"+attr+"` is not available on instance class "+this+".";
        }
    };
    BASE_CLASS.prototype.ClassApply = function(attr){
        var fn = this.ClassAttr(attr);
        if (typeof(fn) !='function'){
            throw "`"+attr+"` is not a function on instance class "+this+".";
        }
        return fn.apply(this, Array.prototype.slice.call(arguments,1)[0]);
    };
    BASE_CLASS.prototype.ClassCall = function(attr){
        var fn = this.ClassAttr(attr);
        if (typeof(fn) !='function'){
            throw "`"+attr+"` is not a function on instance class"+this+".";
        }
        return fn.apply(this, Array.prototype.slice.call(arguments,1));
    };

    // How all classes shall be defined
    globals.Class = function(type){
        /** Create the class **/
        var ret = function(){ return BASE_CLASS.apply(this,arguments); };
        var propidx = 1;
        var propname;

        /** Make sure our type is sane **/
        if (                !type || !type.prototype 
                            || !(type.prototype instanceof BASE_CLASS)){
            type = BASE_CLASS;
            propidx = 0;
        }

        ret.prototype = new type({__NO_INIT__:true});
        ret.prototype._superclass = type;
        ret.prototype._class = ret;

        /** Process all the mix-ins **/
        for (propidx=propidx;propidx<arguments.length;++propidx){

            for (propname in arguments[propidx]){
                ret.prototype[propname] = arguments[propidx][propname];
            }
            
        }
        return ret;
    };

    /* 
     * Instance() is a grammatical alternative to visually make more sense
     * when you need to do something like make an instance out of a bunch
     * of mixins, for example you might have some evented mixin.
     * You can also just use it instead of `new` if you want for some reason.
     */
    globals.Instance = function(){
        return new (globals.Class.apply(globals,arguments))();
    };


})(this);






/* Samples & tests.
 * To perform, call a test, call it.
 * (note the commented calls at the end.)
 */
(function(globals){
    "use strict";
    console = globals.console;
    if (!globals.console){
        var elBody = document.getElementsByTagName('body')[0]
        var console = {
            log: function(){
                var line = '<div>';
                for (var i=0; i<arguments.length; ++i){
                    line += (''+arguments[i]).replace('\n', '<br />')+'&nbsp;&nbsp;&nbsp;';
                }
                line+='</div>';

                elBody.innerHTML = elBody.innerHTML + line;

            }
        };
    } 

    function test_inheritance_chain(){
        // Sample usage
        var organism = Class();

        // Longhand way of defining class properties
        organism.prototype.identify = function(){ return this.name };
        organism.prototype.msg = function(msg){ 
            console.log(this.identify(), msg); return msg; 
        };

        // A subclass of organism
        var animal = Class(organism,{
            // Shorthand way of defining class properties
            walk:function(){ 
                return this.msg("Moving Forward"); 
            }
        })

        // An instance of animal
        var myPet = new animal({name : 'Dirk'});

        // More subclasses & instances
        var dog = Class(animal, {walk: function(){ 
            this.msg("Initializing four legs"); 
            return animal.prototype['walk'].apply(this,arguments)
        }})

        var myPetDog = new dog({name:'Kipper'});

        var human = Class(animal, {
            walk: function(){ 
                this.msg("Upright on two feet"); 
                return animal.prototype['walk'].apply(this,arguments)
            } ,
            walkdog: function(){ 
                if (!this['dog']){
                    this.msg("Wishing to have a dog.");
                    return 0;
                }
                this.msg("Calling out to dog, "+this.dog.identify());
                this.msg("Invoking dog leash");
                this.walk();
                this.dog.walk();
            }
        });
        var myPetHuman = new human({name:"Freddy"});

        // Tests
        console.log(" --- myPet ---");
        console.log(myPet.identify());
        myPet.walk();

        console.log(" --- myPetDog --- ");
        console.log(myPetDog.identify());
        myPetDog.walk();

        console.log(" --- myPetHuman ---");
        console.log(myPetHuman.identify());
        myPetHuman.walk();
        myPetHuman.walkdog();

        console.log(" --- giving myPetHuman a dog to walk --- ");
        myPetHuman.dog = myPetDog;
        myPetHuman.walkdog();


        console.log( " --- Instance comparisons --- \n",
            myPet instanceof animal, "true\n", 
            myPetDog instanceof animal, "true\n",
            myPetDog instanceof dog, "true\n",
            myPetHuman instanceof organism, "true\n",
            myPetDog instanceof human, "false\n"
        );
        
    }

    function test_smart_typing(){
        var targetInst;
        var MixinA = {
            A:'The Letter A'
        }
        var MixinB = {
            B:'The Letter B'
        }

        var EmptyClass = Class({empty:"I'm empty"});

        var EmptySubclass = Class(EmptyClass)
        targetInst = new EmptySubclass();

        console.log("EmptySubclass instance is EmptyClass (true)", targetInst instanceof EmptyClass);

        var EmptySubclassWithMixin = Class(EmptyClass, MixinA )
        targetInst = new EmptySubclassWithMixin();
        console.log("instance with one mixin (The Letter A),(undefined)",  targetInst.A, targetInst.B )

        var EmptySubclassWith2Mixins = Class(EmptyClass, MixinA, MixinB)
        targetInst = new EmptySubclassWith2Mixins();
        console.log("instance with one mixin (The Letter A),(The Letter B)",  targetInst.A, targetInst.B )
        console.log("instance is instanceof EmptyClass (true)", targetInst instanceof EmptyClass);

        var UndefClassWith2Mixins = Class(MixinA, MixinB);
        targetInst = new UndefClassWith2Mixins();
        console.log("instance with one mixin (The Letter A),(The Letter B)",  targetInst.A, targetInst.B )
        console.log("instance is instanceof EmptyClass (false)", targetInst instanceof EmptyClass);


    }

    function test_super_shortcuts(){
        var Building = Class({
            buildIt: function(){              
                console.log("Considering building rules", arguments);
                console.log("Building being constructed");
            }
        });

        //var mybuilding = new Building();

        var House = Class(Building, {
            buildIt: function(){
                console.log("Considering house building rules", arguments);
                //Building.prototype.buildIt.apply(this);
                this.SuperCall(Building, 'buildIt', "Super","Method","Called","With","Arguments");
                console.log("... Compare building to building plans")
                

            }
        });

        var Bungalow = Class(House, {
            buildIt: function(){
                console.log("Reviewing bunglow styles", arguments)
                this.SuperApply(House, 'buildIt', arguments);
                console.log("... Appreciate the style.")
                
            }
        })

        var myBungalow = new Bungalow();
        myBungalow.buildIt('original', 'argument list');


    }

    function test_direct_instance(){
        var oioi = Instance();
        var base = Class().prototype._superclass;
        console.log(oioi instanceof base);
    }

    function test_class_access(){
        var base = Class({
            init: function(){
                console.log("I'm the base class init", arguments)
            }
        });

        var inst = new base({
            init: function(){
                console.log("I am a wild rebellious instance who overrides my class init! WHUT!@?", arguments);
                this.ClassApply('init', arguments);
                console.log("Now I'm getting REALLY wild & rebellious and calling my class constructor TWICE! WHAAAAAT!??");
                this.ClassCall('init', 'AND', 'WITH','HARDCODEd','ARGUMENTS!');
                console.log("Somebody will get fired for this!... stupid free-will");                
            }
        })
    }

    test_inheritance_chain();
    test_smart_typing();
    test_super_shortcuts();
    test_direct_instance();
    test_class_access();

})(this);