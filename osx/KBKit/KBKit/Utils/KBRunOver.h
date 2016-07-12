//
//  KBRunOver.h
//  Keybase
//
//  Created by Gabriel on 4/7/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import <Foundation/Foundation.h>



typedef void (^KBRunCompletion)(id output);
typedef void (^KBRunBlock)(id obj, KBRunCompletion completion);

typedef void (^KBRunOverCompletion)(NSArray *outputs);

@interface KBRunOver : NSObject

@property NSEnumerator *enumerator;
@property (copy) KBRunBlock runBlock;
@property (copy) KBRunOverCompletion completion;

- (void)run;

@end
